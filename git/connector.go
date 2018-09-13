package git

// Connector allows to interact with the Git command-line interface.
type Connector interface {
	GetCommitTemplatePath() (path string, global bool, err error)
	SetCommitTemplate(path string) (err error)
	UnsetCommitTemplate() (err error)
}
