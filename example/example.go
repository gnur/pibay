package main

import (
	"fmt"
	"github.com/gnur/pibay"
)

func main() {
	//example for searching for 1 query string
	err, torrents := pibay.Search("go lang piratebay")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, tor := range torrents {
		fmt.Println(tor.Magnetlink)
		fmt.Println(tor.Title)
		fmt.Println(tor.Size)
		fmt.Println(tor.Uploaded)
		fmt.Println(tor.User)
		fmt.Println(tor.Vipuser)
		fmt.Println(tor.Category)
	}

	//example searching multiple query strings at once
	torch := make(chan pibay.Torrent)
	done := make(chan bool)
	go torrentrecevier(torch)
	go pibay.SearchChannel("gentoo", torch, done)
	go pibay.SearchChannel("ubuntu", torch, done)
	go pibay.SearchChannel("sabayon", torch, done)
	go pibay.SearchChannel("arch linux", torch, done)
	for i := 0; i < 4; i += 1 {
		<-done
	}
}

func torrentrecevier(t chan pibay.Torrent) {
	for tor := range t {
		fmt.Println(tor.Title)
	}
}
