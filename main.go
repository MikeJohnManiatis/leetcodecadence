package main

import (
	"cadence/internal/core"
	"fmt"
)

func main() {
	scraper := &core.Scraper{}
	arrayQuestionMap := scraper.ScrapeLeetCodeQuestions("array")

	fmt.Printf("%v\n", arrayQuestionMap)
}
