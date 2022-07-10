// built with help from: https://blog.logrocket.com/web-scraping-with-go-and-colly/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

type review struct {
	Title    string
	Artist   string
	Genre    string
	Label    string
	Reviewed string
	Score    string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("pitchfork.com"),
		// colly.Async(true),
		colly.MaxDepth(1),
	)

	infoCollector := c.Clone()

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		Delay:       2 * time.Second,
		RandomDelay: time.Second,
	})

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
		tmpReview := review{}
		tmpReview.Title = e.ChildText("h1")
		tmpReview.Artist = e.ChildText("div.BaseWrap-sc-TURhJ.BaseText-fFzBQt.SplitScreenContentHeaderArtist-lgjmiI.eTiIvU.ifBumJ.fUDxJr")

		e.ForEach("div.InfoSliceItem-kovQju.fMqnkQ", func(n int, kf *colly.HTMLElement) {
			if n == 0 {
				tmpReview.Genre = kf.ChildText("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn.glrVeB")
			}
			if n == 1 {
				tmpReview.Label = kf.ChildText("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn.glrVeB")
			}
			if n == 2 {
				tmpReview.Reviewed = kf.ChildText("p.BaseWrap-sc-TURhJ.BaseText-fFzBQt.InfoSliceValue-gSTMso.eTiIvU.bsGTGn.glrVeB")
			}
		})

		tmpReview.Score = e.ChildText("div.ScoreCircle-cJwsOz.cChWcX > p")

		js, err := json.MarshalIndent(tmpReview, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))
	})

	for i := 1; i <= 2063; i++ {
		c.Visit(fmt.Sprintf("https://pitchfork.com/reviews/albums/?page=%d", i))
	}
}
