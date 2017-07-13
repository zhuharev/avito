package avito

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	// BaseURL prefix for all urls
	BaseURL = "https://www.avito.ru"
)

// ParseList parse search result
func ParseList(list []byte) (items []Item, err error) {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(list))
	if err != nil {
		return nil, err
	}

	doc.Find(".catalog-list .item").Each(func(i int, s *goquery.Selection) {
		id, errGoquery := s.Attr("id")
		if !errGoquery {
			err = errors.New("Can not get attribute 'id'")
			return
		}

		var iid int
		iid, _ = strconv.Atoi(id)

		a := s.Find("h3.title a")
		title := strings.TrimSpace(a.Text())
		href, errHref := a.Attr("href")
		if !errHref {
			err = errors.New("Can not get attribute 'href'")
			return
		}
		if iid == 0 {
			iid, err = GetIDFromURL(href)
			if err != nil {
				return
			}
		}
		link := BaseURL + href
		//commission := strings.TrimSpace(s.Find(".about__commission").Text())

		// yet another hell
		html := strings.TrimSpace(s.Find(".about").Text())
		price, _ := strconv.Atoi(regexp.MustCompile(`[^\d]`).ReplaceAllString(html, ""))

		publishedString := strings.TrimSpace(s.Find(".date").Text())

		//address := strings.TrimSpace(s.Find(".address").Text())

		items = append(items, Item{ID: iid, Title: title, URL: link, Price: price, PublishedString: publishedString})
	})
	if err != nil {
		return
	}
	return
}

// GetTextFromURL return text from given url
func GetTextFromURL(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	var text = ""
	doc.Find(".item-description-text p").Each(func(_ int, s *goquery.Selection) {
		text = text + strings.TrimSpace(s.Text()) + "\n\n"
	})

	return text, nil
}

// GetText returns full user text
func GetText(data []byte) (string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	var text = ""
	doc.Find(".item-description-text p").Each(func(_ int, s *goquery.Selection) {
		text = text + strings.TrimSpace(s.Text()) + "\n\n"
	})

	return text, nil
}
