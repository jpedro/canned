// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
package canned

import (
	"os"
	"testing"
)

const (
	testFile     = "test.can"
	testPassword = "test"
)

func TestBootstrap(t *testing.T) {
	os.Setenv("CANNED_TEST_FORMATS", "XSQABaTYTZ1cYdLMUl0ioTUIx")
}

func TestInitCan(t *testing.T) {
	_, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}
}

func TestInitCanNoPassword(t *testing.T) {
	_, err := InitCan(testFile, "")
	if err == nil {
		t.Error("This was supposed to fail")
	}
}

func TestOpenWrongFile(t *testing.T) {
	_, err := OpenCan(testFile+"-wrong", testPassword)
	if err == nil {
		t.Error("Expected to fail to open but just did?!")
	}
}

func TestOpenWrongPassword(t *testing.T) {
	_, err := OpenCan(testFile, testPassword+"-wrong")
	if err == nil {
		t.Error("Expected to fail to open but just did?!")
	}
}

func TestOpenWrongFormat(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// log.Println("Recovered in f", r)
		} else {
			t.Error("Expected to fail to open but just did?!")
		}
	}()

	OpenCan("can.go", testPassword)
}

func TestItemCrud(t *testing.T) {
	name := "name"
	value := "value"
	tag := "testing"

	can, err := OpenCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}
	item, err := can.GetItem(name)
	if err != nil {
		t.Error(err)
	}

	err = can.AddTag(name, tag)
	if err != nil {
		t.Error(err)
	}

	err = can.DelTag(name, tag)
	if err != nil {
		t.Error(err)
	}

	if item.Content != value {
		t.Error("Expected", value, "got", item.Content)
	}
}

func TestCanFormats(t *testing.T) {
	name := "name"
	value := "value"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}

	err = can.Save()
	if err != nil {
		t.Error(err)
	}
}

func TestRenameItem(t *testing.T) {
	var item *Item

	oldName := "oldName"
	newName := "newName"
	value := "value"

	can, err := OpenCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(oldName, value)
	if err != nil {
		t.Error(err)
	}

	item, err = can.GetItem(oldName)
	if err != nil {
		t.Error(err)
	}

	if item.Content != value {
		t.Error("Expected", value, "got", item.Content)
	}

	err = can.RenameItem(oldName, newName)
	if err != nil {
		t.Error(err)
	}

	item, err = can.GetItem(newName)
	if err != nil {
		t.Error(err)
	}

	if item.Content != value {
		t.Error("Expected", value, "got", item.Content)
	}
}

func TestDelItem(t *testing.T) {
	name := "name"
	value := "value"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}

	err = can.DelItem(name)
	if err != nil {
		t.Error(err)
	}

	items := len(can.Items)
	if items > 0 {
		t.Error("Expected to have 0 items got", items)
	}
}

func TestDelWrongItem(t *testing.T) {
	name := "name"
	value := "value"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}

	err = can.DelItem(name + "-wrong")
	if err == nil {
		t.Error("Expected to error but didn't")
	}
}

func TestTagWrongItem(t *testing.T) {
	name := "name"
	value := "value"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.AddTag(name, value)
	if err == nil {
		t.Error("Expected to fail but didn't")
	}
}

func TestTagTwice(t *testing.T) {
	name := "name"
	value := "value"
	tag := "tag"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}

	err = can.AddTag(name, tag)
	if err != nil {
		t.Error(err)
	}

	err = can.AddTag(name, tag)
	if err != nil {
		t.Error(err)
	}

	if len(can.Items[name].Tags) > 1 {
		t.Error("Expected to have only one tag")
	}
}

func TestDelTagNone(t *testing.T) {
	name := "name"
	tag := "tag"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.DelTag(name, tag)
	if err == nil {
		t.Error("Expected to fail")
	}
}

func TestDelTagNoItem(t *testing.T) {
	name := "name"
	value := "value"
	tag := "tag"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		t.Error(err)
	}

	err = can.DelTag(name, tag)
	if err == nil {
		t.Error("Expected to fail")
	}
}

func TestRenameNoItem(t *testing.T) {
	oldName := "oldName"
	newName := "newName"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.RenameItem(oldName, newName)
	if err == nil {
		t.Error("Expected to fail")
	}
}

func TestNoNameItem(t *testing.T) {
	name := ""
	value := "value"

	can, err := InitCan(testFile, testPassword)
	if err != nil {
		t.Error(err)
	}

	err = can.SetItem(name, value)
	if err == nil {
		t.Error("Expected to fail")
	}
}

func TestMoreCoverage(t *testing.T) {
	remove([]string{}, "fail")
	env("BLAH", "I hope this fallback")
	align("", 9)
}
