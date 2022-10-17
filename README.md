go get github.com/marycka9/go-easypronunciation-api

For using:

	c := client.NewClient("your_api_key")

	res, err := c.PhoneticTranslator("en", "You read our book!", true)
	if err != nil {
		log.Println(err)
	}
