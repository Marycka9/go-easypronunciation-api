package main

import (
	"github.com/marycka9/go-easypronunciation-api/client"
	"log"
)

func main() {
	c := client.NewClient("your_api_key")

	res, err := c.PhoneticTranslator("en", "You read our book!", true)
	if err != nil {
		log.Println(err)
	}

	_ = res
}
