// Code generated by "file2go -alias ThemeJuri templates/theme-juri.css"; DO NOT EDIT.
package templates

import "github.com/pscn/file2go/decode"
const contentThemeJuri = `H4sIGAAAAAAA/2RHaGxiV1V0YW5WeWFTNWpjM009AFpXNWpiMlJsWkNCaWVTQm1hV3hsTW1kdgA0ykEKgzAQBdC9p/iQnRBTo5IQT2MmUxSUERPIovTupdRuHy9cIgWvBjAtSM6NE56XHFhLOXMwptba1SHTKrLnjuQw+xZNHXRZ+WB9ceooZ7SmAbSmPkDFhVzk+Qc2QJHlnuwNQ4ByiZKfbhgDlH8s0f/H9B3Oj87OzfsTAAD//66raUahAAAA`

var fileThemeJuri *decode.File

func init() {
	var err error
	fileThemeJuri, err = decode.Init(contentThemeJuri)
	if err != nil {
		panic(err)
	}
}

func ThemeJuriContent() string { return fileThemeJuri.Content() }

func ThemeJuriName() string { return fileThemeJuri.Name() }

func ThemeJuriComment() string { return fileThemeJuri.Comment() }

//eof
