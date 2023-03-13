package infraestructure

import "example.com/greetings/types"

var SecondCatalog = []types.ExcelOptions{
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
				{
					FieldName: "SecondLvl",
					ArrayItem: types.ArrayOptions{
						SubStructItem: []types.ExcelOptions{
							{
								HeaderName: "FirstName Secondlvl",
								FieldName:  "FirstName",
							},
							{
								HeaderName: "LastName Secondlvl",
								FieldName:  "LastName",
							},
							{
								HeaderName: "BirthDate Secondlvl",
								FieldName:  "BirthDate",
							},
						},
					},
				},
			},
		},
	},
	{
		HeaderName: "CreatedAt Header",
		FieldName:  "CreatedAt",
	},
	{
		FieldName: "Demo",
		ArrayItem: types.ArrayOptions{
			SubStructItem: []types.ExcelOptions{
				{
					HeaderName: "FirstName demo",
					FieldName:  "FirstName",
				},
				{
					HeaderName: "LastName demo",
					FieldName:  "LastName",
				},
			},
		},
	},
}
