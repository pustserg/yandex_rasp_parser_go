package main

import "testing"

func TestCity_FullUrl(t *testing.T) {
	city := City{"Moscow", "213"}
	url := city.FullUrl()
	if url != "https://rasp.yandex.ru/city/213" {
		t.Error("Wrong full url calc")
	}
}
