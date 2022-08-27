package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://udemy.com",
		"http://youtube.com",
		"http://vk.com",
	}

	c := make(chan string)

	for _, link := range links {
		go pokeLink(link, c)
	}

	for link := range c {
		go func(link string) {
			time.Sleep(time.Second * 3)
			pokeLink(link, c)
		}(link)
	}
}

func pokeLink(link string, c chan string) {
	response, er := http.Get(link)
	if er != nil {
		fmt.Println(er.Error())
		c <- link
		return
	}
	fmt.Println(link, response.StatusCode)
	c <- link
}
