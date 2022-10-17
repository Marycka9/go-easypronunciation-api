package main

import (
	"github.com/marycka9/go-easypronunciation-api/client"
	"log"
)

func main() {
	c := client.NewClient("your_api_key")

	params := map[string]string{
		"english_phonetics_algorithm": "british_miscellaneous_sources",
	}

	res, err := c.PhoneticTranslator("en", "You read our book!", true, params)
	if err != nil {
		log.Println(err)
	}

	_ = res
}
