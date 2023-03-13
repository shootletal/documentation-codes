package types

type EmailNotify struct {
	CountPending  int
	CountReject   int
	CountApproved int
	FileUrl       string
}
