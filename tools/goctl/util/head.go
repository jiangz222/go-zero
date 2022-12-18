package util

const (
	// DoNotEditHead added to the beginning of a file to prompt the user not to edit
	DoNotEditHead = "// Code generated by goctl. DO NOT EDIT."

	headTemplate = DoNotEditHead + `
// Source: {{.source}}`
)

// GetHead returns a code head string with source filename
func GetHead(source string) string {
	buffer, _ := With("head").Parse(headTemplate).Execute(map[string]interface{}{
		"source": source,
	})
	return buffer.String()
}
