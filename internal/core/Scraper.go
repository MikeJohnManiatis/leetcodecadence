package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

type Scraper struct {
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func startChromeInstance() selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()
	return wd
}

func waitForDynamicRendering(webDriver selenium.WebDriver) {
	for {
		//This element wont be loaded until dynamic rendering is complete.
		_, err := webDriver.FindElement(selenium.ByCSSSelector, ".table.table__XKyc")

		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func gatherQuestionsOnPageByDifficulty(wd selenium.WebDriver, difficulty string) []string {
	var questionList []string

	problemListElement, err := wd.FindElements(selenium.ByCSSSelector, ".reactable-data tr")
	handleError(err)

	for _, element := range problemListElement {

		text, err := element.Text()
		handleError(err)

		anchorElement, err := element.FindElement(selenium.ByCSSSelector, "a")
		handleError(err)

		href, err := anchorElement.GetAttribute("href")
		handleError(err)

		if strings.Contains(text, difficulty) {
			questionList = append(questionList, fmt.Sprintf("https://leetcode.com%s", href))
		}
	}
	return questionList
}

func (scraper *Scraper) ScrapeLeetCodeQuestions(tag string) map[string][]string {
	fmt.Println("Scraping..")

	caps := selenium.Capabilities{"browserName": "safari"}

	wd, err := selenium.NewRemote(caps, "")
	handleError(err)

	defer wd.Quit()

	url := fmt.Sprintf("https://leetcode.com/tag/%s", tag)
	wd.Get(url)

	waitForDynamicRendering(wd)

	questionMap := make(map[string][]string)

	questionMap["Easy"] = gatherQuestionsOnPageByDifficulty(wd, "Easy")
	questionMap["Medium"] = gatherQuestionsOnPageByDifficulty(wd, "Medium")
	questionMap["Hard"] = gatherQuestionsOnPageByDifficulty(wd, "Hard")
	return questionMap
}
