package main

import (
	"encoding/json"
	"example.com/greetings/types"
	"fmt"
	dynamo_types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"reflect"
	"strings"
	"time"
)

var branches = []types.BranchesDynamo{
	{
		Currency: "currency data",
		Name:     "firstName",
		BranchId: "ID",
		LegalRepresentative: []types.PersonLegal{
			{
				FirstName: "first Legal name",
				LastName:  "first legal last",
				SecondLvl: []types.AuxStruct{
					{
						FirstName: "first name second lvl",
						LastName:  "first last second lvl",
					},
					{
						BirthDate: "second birth second lvl",
					},
					{
						LastName: "third last second lvl",
					},
				},
			},
			{
				FirstName: "Second Legal name",
				LastName:  "Second legal last",
				BirthDate: "Second birth",
			},
			{
				LastName: "third legal last",
			},
		},
		ListWC1: []string{"list1", "list2", "list3"},
		MCC:     4,
	},
	{
		Currency: "currency second",
		Name:     "SecondName",
		BranchId: "ID second",
		Shareholder: types.PersonLegal{
			FirstName: "Name test share",
			LastName:  "last shareholder",
		},
		KushkiStatus: "APPROVED",
	},
	{
		Currency:  "currency third",
		Name:      "thirdName",
		BranchId:  "ID third",
		CreatedAt: 1234,
		LegalRepresentative: []types.PersonLegal{
			{
				FirstName: "Cris",
				LastName:  "shoot",
				BirthDate: "birth",
			},
		},
	},
}

type TestStruct struct {
	Responses []interface{}
}

type Client struct {
	Name string
}

type Admin struct {
	Age int
}

type ArrayHelpert struct {
	lastPosition int
}

func main() {
	type KeyGroupBranches struct {
		CustomerID string
		Email      string
	}

	type Aux struct {
		value      string
		Email      string
		CustomerID string
	}

	//tests := []Aux{
	//	{value: "branch1", Email: "email1", CustomerID: "cs1"},
	//	{value: "branch2", Email: "email12", CustomerID: "cs1"},
	//	{value: "branch3", Email: "email1", CustomerID: "cs2"},
	//	{value: "branch4", Email: "email1", CustomerID: "cs1"},
	//}
	//groups := make(map[KeyGroupBranches]map[string][]Aux)
	//
	//for _, branch := range tests {
	//	key := KeyGroupBranches{CustomerID: branch.CustomerID, Email: branch.Email}
	//	secondKey := "CREATE"
	//	if groups[key] == nil {
	//		groups[key] = make(map[string][]Aux)
	//	}
	//	if groups[key][secondKey] == nil {
	//		groups[key][secondKey] = make([]Aux, 0)
	//	}
	//	groups[key][secondKey] = append(groups[key]["CREATE"], branch)
	//}
	//
	//for key, item := range groups {
	//	fmt.Println("key, 世界", key)
	//	fmt.Println("item 世界", item["CREATE"])
	//}

	//var client Client
	//var admin Admin
	//
	//testStruct := TestStruct{
	//	Responses: []interface{}{&client, &admin},
	//}
	//
	//fmt.Println("Hello, 世界", testStruct)
	//array := []int{1, 2, 3}
	//str := "final"
	//sep := "-"
	//arr := reflect.ValueOf(array)
	//concatenated := ConcatArrayAndStringWithSeparator(arr.Interface(), str, sep)
	//fmt.Println(concatenated)

	//var valuesEmail types.EmailNotify
	//request := types.ExcelRequest{
	//	Data:    branches,
	//	Catalog: infraestructure.ThirdCatalog,
	//	AdditionalOptions: types.AdditionalOptions{
	//		ExtraFunc: countFunc,
	//		Params:    []interface{}{&valuesEmail},
	//	},
	//	SheetName: "Branches",
	//}
	//GenerateExcel(request)
	//
	//aux := make(map[string]int)
	//fmt.Println(aux["lol"])

	//letter, _ := excelize.ColumnNumberToName(1)
	//num, _ := excelize.ColumnNameToNumber("A")

	//testCatalog := []types.ExcelOptions{
	//	{
	//		HeaderName: "1",
	//	},
	//	{
	//		HeaderName: "2",
	//	},
	//	{
	//		HeaderName: "3",
	//	},
	//	{
	//		HeaderName: "4",
	//	},
	//	{
	//		HeaderName: "5",
	//	},
	//}
	//
	//testMap := map[string]types.ExcelOptions{
	//	"s1": {
	//		HeaderName: "1",
	//	},
	//	"s2": {
	//		HeaderName: "2",
	//	},
	//	"s3": {
	//		HeaderName: "3",
	//	},
	//	"s4": {
	//		HeaderName: "4",
	//	},
	//	"s5": {
	//		HeaderName: "5",
	//	},
	//	"s6": {
	//		HeaderName: "6",
	//	},
	//}

	//for i := 0; i < len(testCatalog); i++ {
	//	fmt.Println("HeaderName, 世界", testCatalog[i].HeaderName)
	//}
	//
	//for key, _ := range testCatalog {
	//	fmt.Println("Range HeaderName, 世界", key)
	//}
	//
	//for _, j := range testMap {
	//	fmt.Println("testMap Range HeaderName, 世界", j.HeaderName)
	//}

	//fmt.Println("Num, 世界", num)

	//aux := make(map[string]ArrayHelpert)
	//testValue(aux)
	//fmt.Println("test, 世界", aux["test"].lastPosition)
	start, end := GetRangeDateToday("UTC")
	fmt.Println("Hello, 世界 ", start.UnixMilli())
	fmt.Println("Hello, 世界 ", end.UnixMilli())
}

func ConcatArrayAndStringWithSeparator(arr interface{}, str string, sep string) interface{} {
	// Convertimos el array en un slice reflect.Value
	arrValue := reflect.ValueOf(arr)
	if arrValue.Len() > 0 {
		aux := arrValue.Index(0).Interface()
		// Obtenemos el tipo de dato del array
		arrayType := reflect.TypeOf(aux)
		// Creamos un slice vacío con ese tipo de dato
		slice := reflect.MakeSlice(reflect.SliceOf(arrayType), 0, 0)

		// Iteramos sobre los elementos del array
		for i := 0; i < arrValue.Len(); i++ {
			// Obtenemos el valor del elemento
			elemValue := arrValue.Index(i)
			// Añadimos el elemento al slice
			slice = reflect.Append(slice, elemValue)
		}
		// Convertimos el slice en un slice de string
		strSlice := make([]string, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			strSlice[i] = fmt.Sprintf("%v", slice.Index(i))
		}
		// Concatenamos los elementos del slice de string con el separador
		concatenatedStr := strings.Join(strSlice, sep)
		// Concatenamos la variable string al final del resultado
		concatenated := fmt.Sprintf("%s%s", concatenatedStr, str)
		// Retornamos el resultado como un interface{}

		return interface{}(concatenated)
	}
	return ""
}

func testValue(auxFunc map[string]ArrayHelpert) {
	auxFunc["test"] = ArrayHelpert{lastPosition: auxFunc["test"].lastPosition + 1}
}

func buildRequest() types.TokenRequest {
	input := map[string]dynamo_types.KeysAndAttributes{
		"table-name-1": {
			Keys: []map[string]dynamo_types.AttributeValue{
				{
					"field-1": &dynamo_types.AttributeValueMemberS{Value: "aux"},
				},
				{
					"field-1": &dynamo_types.AttributeValueMemberS{Value: "new value"},
				},
			},
		},
		"table-name-2": {
			Keys: []map[string]dynamo_types.AttributeValue{
				{
					"field-1": &dynamo_types.AttributeValueMemberS{Value: "alsdlaslda"},
				},
			},
		},
	}

	listTables := getTableNames(input)
	fmt.Println(listTables)

	return types.TokenRequest{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: types.AuthParametersRequest{
			Username: "backoffice",
			Password: "Zapote1234567!asas",
		},
		ClientId: "3md22mgd5ludtbds2gjisgj3ru",
	}
}

func countFunc(params ...interface{}) {
	branch := params[0].(*types.BranchesDynamo)
	aux := params[1].([]interface{})
	body := aux[0].(*types.EmailNotify)
	switch branch.KushkiStatus {
	case "":
	case "PENDING":
		body.CountPending++
	case "APPROVED":
		body.CountApproved++
	case "REJECTED":
		body.CountReject++
	}
}

func finalBuild() string {
	aux := buildRequest()

	// convert struct to json string
	jsonBytes, _ := json.Marshal(aux)

	return string(jsonBytes)
}

// AnyGetValue get values from data.
func AnyGetValue(s interface{}, fieldNamedStruct string, defaultValue interface{}) interface{} {
	value := reflect.ValueOf(s)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	fieldValue := value.FieldByName(fieldNamedStruct)
	if !fieldValue.IsValid() {
		return defaultValue
	}

	return fieldValue.Interface()
}

func getTableNames(requestItems map[string]dynamo_types.KeysAndAttributes) []string {
	var tableNames []string
	for tableName := range requestItems {
		tableNames = append(tableNames, tableName)
	}
	return tableNames
}

func GetTimeWithTimeZone(timeZone string) time.Time {
	loc, _ := time.LoadLocation(timeZone)

	now := time.Now().In(loc)

	return now
}

// GetRangeDateToday get range dates from today.
func GetRangeDateToday(timeZone string) (time.Time, time.Time) {
	loc, _ := time.LoadLocation(timeZone)
	now := time.Now().In(loc)
	yesterday := now.AddDate(0, 0, -1)
	startDay := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, loc)
	endDay := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	return startDay, endDay
}
