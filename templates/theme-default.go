// Code generated by "file2go -alias ThemeDefault templates/theme-default.css"; DO NOT EDIT.
package templates

import "github.com/pscn/file2go/decode"
const contentThemeDefault = `H4sIGAAAAAAA/2RHaGxiV1V0WkdWbVlYVnNkQzVqYzNNPQBaVzVqYjJSbFpDQmllU0JtYVd4bE1tZHYALMtBqsMgFIXheVZxwFlALy++dHAzaqBdSI3FQOQGFRyU7r3Yev7ZB4eTSMFrAGiEk3P3G55JIkIpZ2aiWqupNrsgcmTjJNKxP6haXYKPXie/GZczRhoArd0fQ12/W34wMdR8a3WwDHWfWh3+GeqytjrM7bK2luH9CQAA//8l5dEXoQAAAA==`

var fileThemeDefault *decode.File

func init() {
	var err error
	fileThemeDefault, err = decode.Init(contentThemeDefault)
	if err != nil {
		panic(err)
	}
}

func ThemeDefaultContent() string { return fileThemeDefault.Content() }

func ThemeDefaultName() string { return fileThemeDefault.Name() }

func ThemeDefaultComment() string { return fileThemeDefault.Comment() }

//eof
