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
	"testing"
)

func TestInitCan(t *testing.T) {
	password := "test"
	file := "test.can"

	_, err := InitCan(file, password)
	if err != nil {
		panic(err)
	}

	// if decrypted != text {
	// 	t.Error("Expected", text, "got", decrypted)
	// }
}

func TestOpenCan(t *testing.T) {
	password := "test"
	file := "test.can"
	name := "test"
	value := "test"

	can, err := OpenCan(file, password)
	if err != nil {
		panic(err)
	}

	err = can.SetItem(name, value)
	if err != nil {
		panic(err)
	}
	item, err := can.GetItem(name)
	if err != nil {
		panic(err)
	}

	if item.Content != "test" {
		t.Error("Expected", value, "got", item.Content)
	}
}
