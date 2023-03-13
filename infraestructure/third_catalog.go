package infraestructure

import "example.com/greetings/types"

var ThirdCatalog = []types.ExcelOptions{
	{
		HeaderName: "Currency Header",
		FieldName:  "Currency",
	},
	{
		HeaderName: "Name Header",
		FieldName:  "Name",
	},
	{
		HeaderName: "BranchId Header",
		FieldName:  "BranchId",
	},
	{
		FieldName: "Shareholder",
		SubStructItem: []types.ExcelOptions{
			{
				HeaderName: "FirstName Header",
				FieldName:  "FirstName",
			},
			{
				HeaderName: "LastName Header",
				FieldName:  "LastName",
			},
			{
				HeaderName: "BirthDate Share",
				FieldName:  "BirthDate",
			},
		},
	},
	{
		FieldName: "LegalRepresentative",
		ArrayItem: types.ArrayOptions{
			SubStructItem: []types.ExcelOptions{
				{
					HeaderName: "FirstName legal",
					FieldName:  "FirstName",
				},
				{
					HeaderName: "LastName legal",
					FieldName:  "LastName",
				},
				{
					HeaderName: "BirthDate legal",
					FieldName:  "BirthDate",
				},
			},
		},
	},
	{
		HeaderName: "CreatedAt Header",
		FieldName:  "CreatedAt",
	},
	{
		HeaderName: "MCC",
		FieldName:  "MCC",
	},
	{
		FieldName:  "ListWC1",
		HeaderName: "LIST WC1",
		ArrayItem: types.ArrayOptions{
			WriteSameCell: true,
			Separator:     ",",
		},
	},
}
