package main

import (
	// "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	out := []string{
		"Apple",
		"BOOK",
		"Book",
		"Go",
		"Hauptbahnhof",
		"January",
		"January",
		"Napkin"}
	in, err := ReadFile("test1.txt")
	if err != nil || !reflect.DeepEqual(out, in) {
		t.Error("wrong answer")
	}
	_, err = ReadFile("lolkek")
	if err == nil {
		t.Error("wrong answer")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	flags := Flags{
		F: true,
		U: false,
		R: false,
		N: false,
		K: 1,
	}
	in := []string{"aaa aA", "a90a bB", "a Aa", "a90DAFa cXCS", "a90SDFa bB"}
	out := []string{"aaa aA", "a90a bB", "a90DAFa cXCS"}
	if !reflect.DeepEqual(RemoveDuplicates(in, flags), out) {
		t.Error("wrong answer")
	}
	flags = Flags{
		F: false,
		U: false,
		R: false,
		N: false,
		K: -1,
	}
	if !reflect.DeepEqual(RemoveDuplicates(in, flags), in) {
		t.Error("wrong answer")
	}
}

func TestMySort(t *testing.T) {
	flags := Flags{
		F: false,
		U: true,
		R: true,
		N: true,
		K: 1,
	}
	in := []string{
		"Apple 3",
		"BOOK 4",
		"Book 5",
		"Go 6",
		"Hauptbahnhof 10",
		"January 7",
		"January 4",
		"Napkin 1"}

	out := []string{
		"Hauptbahnhof 10",
		"January 7",
		"Go 6",
		"Book 5",
		"BOOK 4",
		"Apple 3",
		"Napkin 1"}
	if !reflect.DeepEqual(MySort(in, flags), out) {
		t.Error("wrong answer")
	}
	flags = Flags{
		F: true,
		U: true,
		R: false,
		N: false,
		K: 0,
	}
	out = []string{
		"Apple 3",
		"BOOK 4",
		"Go 6",
		"Hauptbahnhof 10",
		"January 7",
		"Napkin 1"}
	if !reflect.DeepEqual(MySort(in, flags), out) {
		t.Error("wrong answer")
	}
}
