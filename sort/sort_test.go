package main

import (
	// "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	in := []string{"programName", "-o", "kek", "-u", "-k", "3", "-f", "lol"}
	input, output, flags := Parse(in)
	if input != "lol" || output != "kek" {
		t.Error("wrong answer")
	}
	if len(flags) != 3 {
		t.Error("wrong answer")
	}
	if param, has := flags["-k"]; !has || param != 3 {
		t.Error("wrong answer")
	}
	if param, has := flags["-u"]; !has || param != 0 {
		t.Error("wrong answer")
	}
	if param, has := flags["-f"]; !has || param != 0 {
		t.Error("wrong answer")
	}
	in = []string{"programName", "-n", "-r", "azaza"}
	input, output, flags = Parse(in)
	if input != "azaza" || output != "" {
		t.Error("wrong answer")
	}
	if len(flags) != 2 {
		t.Error("wrong answer")
	}
	if param, has := flags["-r"]; !has || param != 0 {
		t.Error("wrong answer")
	}
	if param, has := flags["-n"]; !has || param != 0 {
		t.Error("wrong answer")
	}
}

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
	flags := map[string]int{
		"-f": 0,
		"-k": 1}
	in := []string{"aaa aA", "a90a bB", "a Aa", "a90DAFa cXCS", "a90SDFa bB"}
	out := []string{"aaa aA", "a90a bB", "a90DAFa cXCS"}
	if !reflect.DeepEqual(RemoveDuplicates(in, flags), out) {
		t.Error("wrong answer")
	}
	flags = map[string]int{}
	if !reflect.DeepEqual(RemoveDuplicates(in, flags), in) {
		t.Error("wrong answer")
	}
}

func TestMySort(t *testing.T) {
	flags := map[string]int{
		"-r": 0,
		"-u": 0,
		"-n": 0,
		"-k": 1}
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
	flags = map[string]int{
		"-u": 0,
		"-f": 0,
		"-k": 0}
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
