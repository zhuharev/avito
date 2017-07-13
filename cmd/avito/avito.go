package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ungerik/go-dry"
	"github.com/zhuharev/avito"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("365562083:AAEo1GnCNooSB_ec8GqwdNpV7txLIdbKk_4")
	if err != nil {
		log.Panic(err)
	}

	bts, err := dry.FileGetBytes("https://www.avito.ru/smolensk?q=%D0%B3%D0%B8%D1%80%D0%BE%D1%81%D0%BA%D1%83%D1%82%D0%B5%D1%80")
	if err != nil {
		panic(err)
	}

	items, err := avito.ParseList(bts)
	if err != nil {
		panic(err)
	}
	log.Println(items)
	newIds, err := newIds(ids(items))
	if err != nil {
		panic(err)
	}

	for _, newID := range newIds {
		for _, v := range items {
			if v.ID == newID {
				tpl := `%s
%d
%s

%s

`
				text, err := avito.GetTextFromURL(v.URL)
				if err != nil {
					log.Println(err)
				}
				m := fmt.Sprintf(tpl, v.Title, v.Price, v.PublishedString, text)
				msg := tgbotapi.NewMessage(102710272, m)
				bot.Send(msg)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	oldIDs, err := getSavedIds()
	if err != nil {
		panic(err)
	}
	oldIDs = append(oldIDs, newIds...)
	err = saveIds(oldIDs)
	if err != nil {
		panic(err)
	}
}

func ids(items []avito.Item) []int {
	var ids []int
	for _, v := range items {
		ids = append(ids, v.ID)
	}
	return ids
}

func saveIds(ids []int) error {
	return dry.FileSetJSON("ids", ids)
}

func getSavedIds() ([]int, error) {
	bts, err := ioutil.ReadFile("ids")
	if err != nil {
		return nil, err
	}
	var oldIds []int
	err = json.Unmarshal(bts, &oldIds)
	if err != nil {
		return nil, err
	}
	return oldIds, nil
}

func newIds(ids []int) ([]int, error) {
	var res []int
	oldIds, err := getSavedIds()
	if err != nil {
		return nil, err
	}
	for _, v := range ids {
		if !in(oldIds, v) {
			res = append(res, v)
		}
	}

	return res, nil
}

func in(arr []int, i int) bool {
	for _, v := range arr {
		if v == i {
			return true
		}

	}
	return false
}
