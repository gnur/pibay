package main

import (
	"fmt"
	"github.com/gnur/gopiratebay"
)

func main() {
	err, torrents := gopiratebay.Search("go lang piratebay")
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
}
