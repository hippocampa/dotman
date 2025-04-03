package commons

type StatusString string

const (
	LINKED     StatusString = "linked"
	MISSING    StatusString = "missing"
	CHANGED    StatusString = "changed"
	NOT_LINKED StatusString = "not_linked"
)
