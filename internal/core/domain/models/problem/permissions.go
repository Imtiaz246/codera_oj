package problem

const (
	Author PermitType = "Author"
	Editor PermitType = "Editor"
	Viewer PermitType = "Viewer"
	Tester PermitType = "Tester"
)

type PermitType string

func (t PermitType) IsPermitTypeValid() bool {
	return t == Editor || t == Viewer || t == Tester
}
