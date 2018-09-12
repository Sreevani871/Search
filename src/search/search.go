package search

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	page2KeyWords  = make(map[string]map[string]int)
	pagesInOrder   []string
	keyWord2Weight = make(map[string]int)

	N          = 8
	pageCount  int
	queryCount int
	top        = 5
)

const (
	Page  = "P"
	Query = "Q"
)

type kv struct {
	Key   string
	Value int
}

func AddToPages(pageKeyWords []string) {
	keyWord2Weight = make(map[string]int)
	pageCount++
	if len(pageKeyWords) > N { // If no of keywords exceeds N value stripe off them, only consider N keywords for processing
		pageKeyWords = pageKeyWords[:N-1]
	}
	for i, key := range pageKeyWords {
		keyWord2Weight[strings.ToLower(key)] = N - i
	}
	pageNo := "P" + strconv.Itoa(pageCount)
	page2KeyWords["P"+strconv.Itoa(pageCount)] = keyWord2Weight
	pagesInOrder = append(pagesInOrder, pageNo)
}

func SearchPages(queryKeywords []string) (int, []string) {
	pages := make(map[string]int)
	queryCount++
	var (
		topRelaentPages []string
		last            int
		pagesOrder      []string
	)

	for _, k := range pagesInOrder {
		strength := 0
		pageKeywords := page2KeyWords[k]
		for i, key := range queryKeywords {
			val := pageKeywords[strings.ToLower(key)]
			if val > 0 {
				strength = strength + ((N - i) * val)
			}
		}
		if strength > 0 {
			pages[k] = strength
			pagesOrder = append(pagesOrder, k)

		}
	}

	result := rankByPageStrength(pages, pagesOrder)
	if len(result) < top {
		last = len(result)
	} else {
		last = top
	}
	for _, kv := range result[:last] {
		topRelaentPages = append(topRelaentPages, kv.Key)
	}
	return queryCount, topRelaentPages

}

// Differentiates b/w P and Q cmd and execute accordingly
func Process(args []string) {
	if len(args) > 1 { // Args check to omit P/Q with zero keywords
		cmd := args[0]
		keywords := args[1:]
		switch cmd {
		case Page:
			AddToPages(keywords)
		case Query:
			queryNo, result := SearchPages(keywords)
			fmt.Printf("Q%d:%s\n", queryNo, strings.Join(result, " "))
		default:
			fmt.Println("Enter valid input")
		}
	} else {
		fmt.Println("Please enter input in one of following order\n\t1. P KW1 KW2 KW3...KW8\n\t2. Q KW1 KW2 KW3...KW8")
	}
}

func rankByPageStrength(pages map[string]int, pageOrder []string) []kv {
	var ss []kv
	for _, k := range pageOrder {
		ss = append(ss, kv{k, pages[k]})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss

}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
