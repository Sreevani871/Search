package search

import (
	"testing"
)

type queryTest struct {
	keyWords []string
	Result   []string
}

func TestSearch(t *testing.T) {
	var inputList = [][]string{
		{"Ford", "car", "Review"},
		{"Review", "Car"},
		{"Review", "ford"},
		{"Toyota", "Car"},
		{"Honda", "Car"},
		{"Car"},
	}
	var queryList = []queryTest{
		{[]string{"Ford"}, []string{"P1", "P3"}},
		{[]string{"Car"}, []string{"P6", "P1", "P2", "P4", "P5"}},
		{[]string{"Review"}, []string{"P2", "P3", "P1"}},
		{[]string{"Ford", "Review"}, []string{"P3", "P1", "P2"}},
		{[]string{"Ford", "Car"}, []string{"P1", "P3", "P6", "P2", "P4"}},
		{[]string{"cooking", "French"}, []string{}},
	}

	t.Log("AddToPages")
	for _, keyWords := range inputList {
		AddToPages(keyWords)
	}

	t.Run("Test SearchPages", func(t *testing.T) {
		for _, query := range queryList {
			_, pages := SearchPages(query.keyWords)
			for i, page := range query.Result {
				if page != pages[i] {
					t.Errorf("expected page:%s but got page:%s", page, pages[i])
					t.FailNow()
				}
			}
		}
	})

}
