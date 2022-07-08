// built with help from: https://blog.logrocket.com/web-scraping-with-go-and-colly/

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type review struct {
	Title    string
	Artist   string
	Genre    string
	Label    string
	Score    string
	CoverArt string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("pitchfork.com"),
		// colly.Async(true),
		colly.MaxDepth(1),
	)

	infoCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting review URL:", r.URL)
	})

	c.OnHTML("div.review", func(e *colly.HTMLElement) {
		reviewUrl := e.ChildAttr(".review__link", "href")
		reviewUrl = e.Request.AbsoluteURL(reviewUrl)
		infoCollector.Visit(reviewUrl)
	})

	infoCollector.OnHTML("article", func(e *colly.HTMLElement) {
		// fmt.Println(e.ChildAttr("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn.glrVeB", "p"))
		fmt.Println(e.ChildText("div.InfoSliceItem-kovQju.fMqnkQ > p"))
		// tmpReview := review{}
		// tmpReview.Title = e.ChildText("h1")
		// tmpReview.Artist = [e.]
		// tmpReview.Genre = e.ChildText("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn")
		// tmpReview.Label = e.ChildText("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn")
		// tmpProfile.Photo = e.ChildAttr("#name-poster", "src")
		// tmpProfile.JobTitle = e.ChildText("#name-job-categories > a > span.itemprop")
		// tmpProfile.BirthDate = e.ChildAttr("#name-born-info time", "datetime")
	})

	// for i := 1; i <= 2063; i++ {
	// 	url := "https://pitchfork.com/reviews/albums/?page=" + strconv.Itoa(int(i))

	// 	c.Visit(url)
	// }

	c.Visit("https://pitchfork.com/reviews/albums/?page=1")
}
