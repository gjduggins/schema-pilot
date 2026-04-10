package db

type ScmAdapter interface {
	FetchFile(filePath, tag string) ([]byte, error)
}
