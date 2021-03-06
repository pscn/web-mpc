// Code generated by "file2go -v -t -o templates/files.go templates/*.html templates/*.js templates/*.css templates/*.ico"; DO NOT EDIT.

// Testing for included files:
// → templates/icons.html
// → templates/index.html
// → templates/javascript.js
// → templates/style.css
// → templates/theme-default.css
// → templates/theme-juri.css
// → templates/favicon.ico

package templates

import (
	"testing"
)

func TestContentDoesNotExist(t *testing.T) {
	_, err := Content("")
	if err == nil {
		t.Fatalf("Content: returned no error")
	}
}

func TestContentExists(t *testing.T) {
  var err error
	_, err = Content("templates/icons.html")
	if err != nil {
		t.Fatalf("Content \"templates/icons.html\" not found: %s", err)
	}
	_, err = Content("templates/index.html")
	if err != nil {
		t.Fatalf("Content \"templates/index.html\" not found: %s", err)
	}
	_, err = Content("templates/javascript.js")
	if err != nil {
		t.Fatalf("Content \"templates/javascript.js\" not found: %s", err)
	}
	_, err = Content("templates/style.css")
	if err != nil {
		t.Fatalf("Content \"templates/style.css\" not found: %s", err)
	}
	_, err = Content("templates/theme-default.css")
	if err != nil {
		t.Fatalf("Content \"templates/theme-default.css\" not found: %s", err)
	}
	_, err = Content("templates/theme-juri.css")
	if err != nil {
		t.Fatalf("Content \"templates/theme-juri.css\" not found: %s", err)
	}
	_, err = Content("templates/favicon.ico")
	if err != nil {
		t.Fatalf("Content \"templates/favicon.ico\" not found: %s", err)
	}
}

func TestContentMustDoesNotExist(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("ContentMust: should have panic'd")
		}
	}()
	_ = ContentMust("")
	t.Fatalf("ContentMust: should have panic'd")
}

func TestContentMustExist(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ContentMust: should not have panic'd")
		}
	}()
	_ = ContentMust("templates/icons.html")
	_ = ContentMust("templates/index.html")
	_ = ContentMust("templates/javascript.js")
	_ = ContentMust("templates/style.css")
	_ = ContentMust("templates/theme-default.css")
	_ = ContentMust("templates/theme-juri.css")
	_ = ContentMust("templates/favicon.ico")
}

// eof
