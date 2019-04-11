// Code generated by "file2go -alias Index templates/index.html"; DO NOT EDIT.
package templates

import (
  "time"

  "github.com/pscn/file2go/decode"
)

const contentIndex = `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`

var fileIndex *decode.File

func init() {
	var err error
	fileIndex, err = decode.Init(contentIndex)
	if err != nil {
		panic(err)
	}
}

func IndexContent() string { return fileIndex.Content() }

func IndexName() string { return fileIndex.Name() }

func IndexModTime() time.Time { return fileIndex.ModTime() }

func IndexComment() string { return fileIndex.Comment() }

//eof
