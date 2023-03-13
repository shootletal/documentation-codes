package types

type ArrayOptions struct {
	WriteSameCell bool
	Separator     string
	SubStructItem []ExcelOptions
}

type ExcelOptions struct {
	ArrayItem     ArrayOptions
	HeaderName    string
	FieldName     string
	SubStructItem []ExcelOptions
}

type AdditionalOptions struct {
	ExtraFunc func(params ...interface{})
	Params    []interface{}
}

type ExcelRequest struct {
	Data              interface{}
	Catalog           []ExcelOptions
	SheetName         string
	AdditionalOptions AdditionalOptions
}
