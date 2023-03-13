package main

import (
	"bytes"
	"example.com/greetings/types"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type ArrayHelper struct {
	isArray        bool
	beforeMaxIndex int
	maxIndex       int
	currentIndex   int
	sumPosition    int
}

type ArraySheet struct {
	isFirst bool
	index   int
}

const (
	rowHeader   = 1
	rowNext     = 2
	prefixArray = "ARRAY_CELL_"
)

func GenerateExcel(request types.ExcelRequest) io.Reader {
	// Create new file excel.
	f := excelize.NewFile()

	// Create helpers to array fields.
	arrayHelper := make(map[string]ArrayHelper)

	// Add new sheet in file.
	index, _ := f.NewSheet(request.SheetName)

	// Generate headers to sheet.
	generateHeader(request.Catalog, f, request.SheetName, 1, ArrayHelper{})

	// Generate lines.
	v := reflect.ValueOf(request.Data)
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Addr().Interface()

		if request.AdditionalOptions.ExtraFunc != nil {
			request.AdditionalOptions.ExtraFunc(item, request.AdditionalOptions.Params)
		}

		generateLines(item, request.Catalog, f, arrayHelper, request.SheetName, i+rowNext, 1)
	}

	// Set active first sheet.
	f.SetActiveSheet(index)

	mergeArrayColumns(f, arrayHelper, request.SheetName)
	f.DeleteSheet("Sheet1")

	f.SaveAs("SecondTest.xlsx")

	buf, _ := f.WriteToBuffer()

	return bytes.NewReader(buf.Bytes())

}

func generateHeader(catalog []types.ExcelOptions, excelFile *excelize.File, sheetName string, colIndex int, arrayHelper ArrayHelper) {
	currentIndex := colIndex
	for _, options := range catalog {
		if !reflect.DeepEqual(options.ArrayItem, types.ArrayOptions{}) && !options.ArrayItem.WriteSameCell {
			letter, _ := excelize.ColumnNumberToName(currentIndex)
			newSheetName := prefixArray + options.FieldName
			excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowHeader), newSheetName)
			_, _ = excelFile.NewSheet(newSheetName)
			generateHeader(options.ArrayItem.SubStructItem, excelFile, newSheetName, colIndex, arrayHelper)
		} else {
			if len(options.SubStructItem) != 0 {
				generateHeader(options.SubStructItem, excelFile, sheetName, currentIndex, arrayHelper)
				currentIndex = currentIndex + len(options.SubStructItem) - 1
			} else {
				headerName := options.HeaderName

				if arrayHelper.isArray {
					headerName = headerName + " " + strconv.Itoa(arrayHelper.maxIndex)
				}

				letter, _ := excelize.ColumnNumberToName(currentIndex)
				excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowHeader), headerName)
			}
		}
		currentIndex++
	}
}

func generateLines(data interface{}, catalog []types.ExcelOptions, excelFile *excelize.File, arrayHelper map[string]ArrayHelper, sheetName string, rowIndex int, colIndex int) {
	v := reflect.ValueOf(data).Elem()
	currentIndex := colIndex

	for _, option := range catalog {
		field := v.FieldByName(option.FieldName)

		if field.IsValid() {
			if !reflect.DeepEqual(option.ArrayItem, types.ArrayOptions{}) {
				if option.ArrayItem.WriteSameCell {
					generateSameLineArray(field.Interface(), option, excelFile, sheetName, rowIndex, currentIndex)
				} else {
					generateArrayLines(field.Interface(), option, excelFile, rowIndex, arrayHelper)
				}

			} else {
				if len(option.SubStructItem) != 0 {
					switch field.Kind() {
					case reflect.Struct:
						generateLines(field.Addr().Interface(), option.SubStructItem, excelFile, arrayHelper, sheetName, rowIndex, currentIndex)
						currentIndex = currentIndex + len(option.SubStructItem) - 1
					case reflect.Slice:
						fmt.Println("array")
					}
				} else {
					letter, _ := excelize.ColumnNumberToName(currentIndex)
					val := field.Interface()
					excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowIndex), val)
				}
			}

		}
		currentIndex++
	}
}

func generateArrayLines(data interface{}, catalogItem types.ExcelOptions, excelFile *excelize.File, rowIndex int, arrayHelper map[string]ArrayHelper) {
	t := reflect.ValueOf(data)
	sheetNameSlice := prefixArray + catalogItem.FieldName

	for j := 0; j < t.Len(); j++ {
		subItem := t.Index(j).Addr().Interface()
		a, ok := arrayHelper[sheetNameSlice]

		if ok {
			if a.maxIndex < t.Len() {
				a.maxIndex = t.Len()
			}
		} else {
			a = ArrayHelper{
				isArray:      true,
				sumPosition:  len(catalogItem.ArrayItem.SubStructItem),
				maxIndex:     1,
				currentIndex: 1,
			}

		}
		generateLines(subItem, catalogItem.ArrayItem.SubStructItem, excelFile, arrayHelper, sheetNameSlice, rowIndex, a.currentIndex)
		a.currentIndex = a.currentIndex + a.sumPosition
		arrayHelper[sheetNameSlice] = a
	}

	a, _ := arrayHelper[sheetNameSlice]

	if a.beforeMaxIndex < a.maxIndex {
		current := 1
		for k := 0; k < a.maxIndex; k++ {
			generateHeader(catalogItem.ArrayItem.SubStructItem, excelFile, sheetNameSlice, current, ArrayHelper{isArray: true, maxIndex: k + 1})
			current = current + a.sumPosition
		}
		a.beforeMaxIndex = a.maxIndex
	}
	a.currentIndex = 1
	arrayHelper[sheetNameSlice] = a
}

func generateSameLineArray(data interface{}, catalog types.ExcelOptions, f *excelize.File, sheetName string, rowIndex int, colIndex int) {
	arrValue := reflect.ValueOf(data)

	if arrValue.Len() > 0 {
		strSlice := make([]string, arrValue.Len())
		for i := 0; i < arrValue.Len(); i++ {
			elemValue := arrValue.Index(i).Interface()
			strSlice[i] = fmt.Sprintf("%v", elemValue)
		}

		strConcat := strings.Join(strSlice, catalog.ArrayItem.Separator)

		letter, _ := excelize.ColumnNumberToName(colIndex)

		_ = f.SetCellValue(sheetName, letter+strconv.Itoa(rowIndex), strConcat)
	}
}

func mergeArrayColumns(f *excelize.File, helper map[string]ArrayHelper, sheetName string) {
	for key, _ := range helper {
		columns, _ := f.GetCols(key)
		primaryRows, _ := f.GetRows(sheetName)
		findColumn := findIndex(primaryRows[0], key)
		position := findColumn + 1
		letterRemove, _ := excelize.ColumnNumberToName(position)
		_ = f.RemoveCol(sheetName, letterRemove)

		for _, column := range columns {
			letter, _ := excelize.ColumnNumberToName(position)
			_ = f.InsertCols(sheetName, letter, 1)
			_ = f.SetSheetCol(sheetName, letter+strconv.Itoa(1), &column)
			position++
		}
		_ = f.DeleteSheet(key)
	}
}

func findIndex(slice interface{}, searchValue interface{}) int {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("indexOf: no se puede buscar en un valor no slice")
	}

	for i := 0; i < s.Len(); i++ {
		value := s.Index(i).Interface()
		if reflect.DeepEqual(searchValue, value) {
			return i
		}
	}

	return -1
}

//
//func GenerateExcel(request types.ExcelRequest) *excelize.File {
//	// Create new file excel.
//	f := excelize.NewFile()
//
//	// Create helpers to array fields.
//	arrayHelper := make(map[string]ArrayHelper)
//
//	// Add new sheet in file.
//	index, _ := f.NewSheet(request.SheetName)
//
//	// Generate headers to sheet.
//	generateHeader(request.Catalog, f, request.SheetName)
//
//	// Generate lines.
//	v := reflect.ValueOf(request.Data)
//	for i := 0; i < v.Len(); i++ {
//		item := v.Index(i).Addr().Interface()
//
//		if request.Additional != nil {
//			request.Additional(item, request.Params)
//		}
//
//		generateLines(item, request.Catalog, f, arrayHelper, request.SheetName, i+rowNext, 1)
//	}
//
//	// Establece la primera hoja como hoja activa.
//	f.SetActiveSheet(index)
//
//	colData, _ := f.GetCols(request.SheetName)
//	f.SetSheetCol(request.SheetName, "H1", &colData[3])
//	f.InsertCols(request.SheetName, "B", 2)
//
//	_ = f.SaveAs("testExcel.xlsx")
//
//	return f
//
//}
//
//func generateHeader(catalog map[string]types.ExcelOptions, excelFile *excelize.File, sheetName string) {
//	for key, options := range catalog {
//		if !reflect.DeepEqual(options.ArrayItem, types.ArrayOptions{}) {
//			newSheetName := prefixArray + key
//			excelFile.SetCellValue(sheetName, options.ArrayItem.InitLetterCell+strconv.Itoa(rowHeader), newSheetName)
//			excelFile.NewSheet(newSheetName)
//		} else {
//			if len(options.SubStructItem) != 0 {
//				generateHeader(options.SubStructItem, excelFile, sheetName)
//			} else {
//				excelFile.SetCellValue(sheetName, options.CellLetter+strconv.Itoa(rowHeader), options.HeaderName)
//			}
//		}
//
//	}
//}
//
//func generateArrayHeader(catalog map[string]types.ExcelOptions, excelFile *excelize.File, sheetName string, options ArrayHelper) {
//
//}
//
//func generateLines(data interface{}, catalog map[string]types.ExcelOptions, excelFile *excelize.File, arrayHelper map[string]ArrayHelper, sheetName string, rowIndex int, colIndex int) {
//	v := reflect.ValueOf(data).Elem()
//
//	for key, option := range catalog {
//		field := v.FieldByName(key)
//
//		if field.IsValid() {
//			if !reflect.DeepEqual(option.ArrayItem, types.ArrayOptions{}) {
//				sheetNameSlice := prefixArray + key
//
//				t := reflect.ValueOf(field.Interface())
//
//				for j := 0; j < t.Len(); j++ {
//					subItem := t.Index(j).Addr().Interface()
//					generateLines(subItem, option.ArrayItem.SubStructItem, excelFile, arrayHelper, sheetNameSlice, rowIndex, arrayHelper[sheetNameSlice].lastPosition)
//				}
//
//			} else {
//				if len(option.SubStructItem) != 0 {
//					switch field.Kind() {
//					case reflect.Struct:
//						generateLines(field.Addr().Interface(), option.SubStructItem, excelFile, arrayHelper, sheetName, rowIndex, colIndex)
//					case reflect.Slice:
//
//					}
//				} else {
//					excelFile.SetCellValue(sheetName, option.CellLetter+strconv.Itoa(rowIndex), field.Interface())
//				}
//			}
//		}
//	}
//
//}
//
//func buildArray(data interface{}, catalog map[string]types.ExcelOptions, excelFile *excelize.File, arrayHelper map[string]ArrayHelper, sheetName string, rowIndex int, colIndex int) {
//	rowsData, _ := excelFile.GetRows(sheetName)
//	excelFile.
//
//}
//
//func buildCells(data interface{}, catalog map[string]types.ExcelOptions, excelFile *excelize.File, sheetName string, rowIndex int) {
//	dataValue := reflect.ValueOf(data).Elem()
//	dataType := dataValue.Type()
//
//	for i := 0; i < dataValue.NumField(); i++ {
//		field := dataValue.Field(i)
//		fieldName := dataType.Field(i).Name
//		catalogOptions := catalog[fieldName]
//
//		if !reflect.DeepEqual(catalogOptions, types.ExcelOptions{}) {
//			switch field.Kind() {
//			case reflect.Struct:
//				buildCells(field.Addr().Interface(), catalogOptions.SubStructItem, excelFile, sheetName, rowIndex)
//			case reflect.Slice:
//				elemType := field.Type().Elem()
//				if elemType.Kind() == reflect.Struct {
//					_ = reflect.New(elemType).Elem()
//				}
//			default:
//				excelFile.SetCellValue(sheetName, catalogOptions.CellLetter+strconv.Itoa(rowIndex), field.Interface())
//			}
//		}
//
//	}
//}

//func GenerateExcel(request types.ExcelRequest) *excelize.File {
//	// Create new file excel.
//	f := excelize.NewFile()
//
//	// Create helpers to array fields.
//	arrayHelper := make(map[string]ArrayHelper)
//
//	// Add new sheet in file.
//	index, _ := f.NewSheet(request.SheetName)
//	_, _ = f.NewSheet("ARRAY_CELL_LegalRepresentative")
//	_, _ = f.NewSheet("ARRAY_CELL_Demo")
//
//	// Generate headers to sheet.
//	generateHeader(request.Catalog, f, request.SheetName, 1, ArrayHelper{})
//
//	// Generate lines.
//	v := reflect.ValueOf(request.Data)
//	for i := 0; i < v.Len(); i++ {
//		item := v.Index(i).Addr().Interface()
//
//		if request.Additional != nil {
//			request.Additional(item, request.Params)
//		}
//
//		generateLines(item, request.Catalog, f, arrayHelper, request.SheetName, i+rowNext, 1, i+1, true)
//	}
//
//	// Establece la primera hoja como hoja activa.
//	f.SetActiveSheet(index)
//
//	//mergeArrayColumns(f, arrayHelper, request.SheetName)
//
//	_ = f.SaveAs("testExcel.xlsx")
//
//	return f
//
//}
//
//func generateHeader(catalog []types.ExcelOptions, excelFile *excelize.File, sheetName string, colIndex int, arrayHelper ArrayHelper) {
//	currentIndex := colIndex
//	for _, options := range catalog {
//		if !reflect.DeepEqual(options.ArrayItem, types.ArrayOptions{}) {
//			newSheetName := prefixArray + options.FieldName
//
//			if arrayHelper.isArray {
//				newSheetName = newSheetName + "_" + strconv.Itoa(arrayHelper.maxIndex)
//			}
//			letter, _ := excelize.ColumnNumberToName(currentIndex)
//
//			excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowHeader), newSheetName)
//		} else {
//			if len(options.SubStructItem) != 0 {
//				generateHeader(options.SubStructItem, excelFile, sheetName, currentIndex, arrayHelper)
//				currentIndex = currentIndex + len(options.SubStructItem) - 1
//			} else {
//				headerName := options.HeaderName
//
//				if arrayHelper.isArray {
//					headerName = headerName + " " + strconv.Itoa(arrayHelper.maxIndex)
//				}
//
//				letter, _ := excelize.ColumnNumberToName(currentIndex)
//				excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowHeader), headerName)
//			}
//		}
//		currentIndex++
//	}
//}
//
//func generateLines(data interface{}, catalog []types.ExcelOptions, excelFile *excelize.File, arrayHelper map[string]ArrayHelper, sheetName string, rowIndex int, colIndex int, index int, isFirstLvl bool) {
//	v := reflect.ValueOf(data).Elem()
//	currentIndex := colIndex
//
//	for _, option := range catalog {
//		field := v.FieldByName(option.FieldName)
//
//		if field.IsValid() {
//			if !reflect.DeepEqual(option.ArrayItem, types.ArrayOptions{}) {
//				fieldName := option.FieldName
//				value := field.Interface()
//				fmt.Println(fieldName)
//				fmt.Println(value)
//				generateArrayLines(field.Interface(), option, excelFile, rowIndex, index, isFirstLvl, arrayHelper)
//			} else {
//				if len(option.SubStructItem) != 0 {
//					switch field.Kind() {
//					case reflect.Struct:
//						generateLines(field.Addr().Interface(), option.SubStructItem, excelFile, arrayHelper, sheetName, rowIndex, currentIndex, index, isFirstLvl)
//						currentIndex = currentIndex + len(option.SubStructItem) - 1
//					case reflect.Slice:
//						fmt.Println("array")
//					}
//				} else {
//					letter, _ := excelize.ColumnNumberToName(currentIndex)
//					val := field.Interface()
//					excelFile.SetCellValue(sheetName, letter+strconv.Itoa(rowIndex), val)
//				}
//			}
//
//		}
//		currentIndex++
//	}
//}
//
//func generateArrayLines(data interface{}, catalogItem types.ExcelOptions, excelFile *excelize.File, rowIndex int, index int, isFirstLvl bool, arrayHelper map[string]ArrayHelper) {
//	t := reflect.ValueOf(data)
//
//	sheetNameSlice := updateSheetByArray(excelFile, catalogItem.FieldName, index, isFirstLvl)
//	a, ok := arrayHelper[sheetNameSlice]
//
//	for j := 0; j < t.Len(); j++ {
//		subItem := t.Index(j).Addr().Interface()
//
//		if ok {
//			if a.maxIndex < t.Len() {
//				a.maxIndex = t.Len()
//			}
//		} else {
//			a = ArrayHelper{
//				isArray:      true,
//				sumPosition:  len(catalogItem.ArrayItem.SubStructItem),
//				maxIndex:     1,
//				currentIndex: 1,
//			}
//
//		}
//		generateLines(subItem, catalogItem.ArrayItem.SubStructItem, excelFile, arrayHelper, sheetNameSlice, rowIndex, a.currentIndex, j+1, false)
//		a.currentIndex = a.currentIndex + a.sumPosition
//		arrayHelper[sheetNameSlice] = a
//	}
//
//	if a.beforeMaxIndex < a.maxIndex {
//		current := 1
//		for k := 0; k < a.maxIndex; k++ {
//			generateHeader(catalogItem.ArrayItem.SubStructItem, excelFile, sheetNameSlice, current, ArrayHelper{isArray: true, maxIndex: k + 1})
//			current = current + a.sumPosition
//		}
//		a.beforeMaxIndex = a.maxIndex
//	}
//	a.currentIndex = 1
//	arrayHelper[sheetNameSlice] = a
//
//}
//
//func updateSheetByArray(f *excelize.File, fieldName string, index int, isFirstLvl bool) string {
//	sheetName := prefixArray + fieldName
//
//	if !isFirstLvl {
//		sheetName = prefixArray + fieldName + strconv.Itoa(index)
//		_, _ = f.NewSheet(sheetName)
//	}
//
//	return sheetName
//}
//
//func mergeArrayColumns(f *excelize.File, helper map[string]ArrayHelper, sheetName string) {
//	for key, _ := range helper {
//		columns, _ := f.GetCols(key)
//		primaryRows, _ := f.GetRows(sheetName)
//		findColumn := findIndex(primaryRows[0], key)
//		position := findColumn + 1
//		letterRemove, _ := excelize.ColumnNumberToName(position)
//		_ = f.RemoveCol(sheetName, letterRemove)
//
//		for _, column := range columns {
//			letter, _ := excelize.ColumnNumberToName(position)
//			_ = f.InsertCols(sheetName, letter, 1)
//			_ = f.SetSheetCol(sheetName, letter+strconv.Itoa(1), &column)
//			position++
//		}
//		_ = f.DeleteSheet(key)
//	}
//}
//
//func findIndex(slice interface{}, searchValue interface{}) int {
//	s := reflect.ValueOf(slice)
//
//	if s.Kind() != reflect.Slice {
//		panic("indexOf: no se puede buscar en un valor no slice")
//	}
//
//	for i := 0; i < s.Len(); i++ {
//		value := s.Index(i).Interface()
//		if reflect.DeepEqual(searchValue, value) {
//			return i
//		}
//	}
//
//	return -1
//}
