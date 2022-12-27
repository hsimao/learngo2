package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	href     string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://www.cakeresume.com/jobs/frontend%20engineer"

func Run() {

	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getDetailByPage(i + 1)

	}
}

func getDetailByPage(page int) {
	pageUrl := baseURL + "?page=" + strconv.Itoa(page)
	fmt.Println("pagpageUrle", pageUrl)

	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("div[class^=JobSearchItem_wrapper__]")
	searchCards.Each(func(i int, card *goquery.Selection) {

		// 內頁連結
		href, _ := card.Find("a[class^=JobSearchItem_jobTitle__]").Attr("href")
		// 標題
		title := card.Find("a[class^=JobSearchItem_jobTitle__]").Text()
		fmt.Println("href", href)
		fmt.Println("title", title)

		// 地點
		location := card.Find("a[class^=JobSearchItem_featureSegmentLink__]").Text()
		fmt.Println("location", location)

		// 薪水, 先找到 dollar icon, 在回頭往上找
		salary := card.Find(".fa-dollar-sign").Closest("div[class^=InlineMessage_inlineMessage__]").Find("div[class^=InlineMessage_label__]").Text()
		fmt.Println("salary ", salary)
		// 簡介
		summary := card.Find("div[class^=JobSearchItem_description__]").Text()
		fmt.Println("summary ", summary)

	})

}

func getPages() int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	checkErr(err)

	// fake browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")

	resp, err := client.Do(req)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	// Pagination_itemNumber__ 開頭的 className
	pages := doc.Find("[class^=Pagination_itemNumber__]").Length()

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func clearString(str string) string {

}
