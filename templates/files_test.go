// Code generated by "file2go -v -t -o templates/files.go dist/*.html dist/*.ico dist/js/*.js dist/js/*.map dist/css/*.css"; DO NOT EDIT.

// Testing for included files:
// → dist/index.html
// → dist/favicon.ico
// → dist/js/app.9e97c486.js
// → dist/js/chunk-vendors.1e34a291.js
// → dist/js/app.9e97c486.js.map
// → dist/js/chunk-vendors.1e34a291.js.map
// → dist/css/app.dde38c96.css

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

	_, err = Content("dist/index.html")
	if err != nil {
		t.Fatalf("Content \"dist/index.html\" not found: %s", err)
	}
	_, err = Content("dist/favicon.ico")
	if err != nil {
		t.Fatalf("Content \"dist/favicon.ico\" not found: %s", err)
	}
	_, err = Content("dist/js/app.9e97c486.js")
	if err != nil {
		t.Fatalf("Content \"dist/js/app.9e97c486.js\" not found: %s", err)
	}
	_, err = Content("dist/js/chunk-vendors.1e34a291.js")
	if err != nil {
		t.Fatalf("Content \"dist/js/chunk-vendors.1e34a291.js\" not found: %s", err)
	}
	_, err = Content("dist/js/app.9e97c486.js.map")
	if err != nil {
		t.Fatalf("Content \"dist/js/app.9e97c486.js.map\" not found: %s", err)
	}
	_, err = Content("dist/js/chunk-vendors.1e34a291.js.map")
	if err != nil {
		t.Fatalf("Content \"dist/js/chunk-vendors.1e34a291.js.map\" not found: %s", err)
	}
	_, err = Content("dist/css/app.dde38c96.css")
	if err != nil {
		t.Fatalf("Content \"dist/css/app.dde38c96.css\" not found: %s", err)
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
	_ = ContentMust("dist/index.html")
	_ = ContentMust("dist/favicon.ico")
	_ = ContentMust("dist/js/app.9e97c486.js")
	_ = ContentMust("dist/js/chunk-vendors.1e34a291.js")
	_ = ContentMust("dist/js/app.9e97c486.js.map")
	_ = ContentMust("dist/js/chunk-vendors.1e34a291.js.map")
	_ = ContentMust("dist/css/app.dde38c96.css")
}

func TestFilenames(t *testing.T) {
	filenames := Filenames()
	i := 0
	if filenames[i] != "dist/index.html" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/index.html")
	}
	i++
	if filenames[i] != "dist/favicon.ico" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/favicon.ico")
	}
	i++
	if filenames[i] != "dist/js/app.9e97c486.js" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/js/app.9e97c486.js")
	}
	i++
	if filenames[i] != "dist/js/chunk-vendors.1e34a291.js" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/js/chunk-vendors.1e34a291.js")
	}
	i++
	if filenames[i] != "dist/js/app.9e97c486.js.map" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/js/app.9e97c486.js.map")
	}
	i++
	if filenames[i] != "dist/js/chunk-vendors.1e34a291.js.map" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/js/chunk-vendors.1e34a291.js.map")
	}
	i++
	if filenames[i] != "dist/css/app.dde38c96.css" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "dist/css/app.dde38c96.css")
	}
	i++
}

// eof
