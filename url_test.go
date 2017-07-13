package avito

import "testing"

func TestURL(t *testing.T) {
	u := "https://www.avito.ru/smolensk/sport_i_otdyh/giroskuter_10_wmotion_wm8_890202998"
	expectedID := 890202998
	id, err := GetIDFromURL(u)
	if err != nil {
		t.Fatal(err)
	}
	if id != expectedID {
		t.Fatal("bad id")
	}
}
