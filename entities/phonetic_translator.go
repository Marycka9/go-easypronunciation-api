package entities

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"net/url"
)

var availableLangs = map[string]string{
	"en": "english",
	"fr": "french",
	"de": "german",
	"it": "italian",
	"ja": "japanese",
	"pt": "portuguese",
	"ru": "russian",
	"es": "spanish",
}

const urlQuery = "https://easypronunciation.com/%s-api.php"

type PhoneticTranslatorRequest struct {
	Phrase   string `json:"phrase"`
	Language string `json:"language"`
	Base64   bool   `json:"base64"`
}

func NewPhoneticTranslatorRequest(phrase, language string, base64 bool) *PhoneticTranslatorRequest {
	if base64 {
		phrase = b64.StdEncoding.EncodeToString([]byte(phrase))
	}
	return &PhoneticTranslatorRequest{
		Phrase:   phrase,
		Language: language,
		Base64:   base64,
	}
}

func (s PhoneticTranslatorRequest) GetParams(token string) url.Values {
	base64 := "0"
	if s.Base64 {
		base64 = "1"
	}
	return url.Values{
		"access_token": []string{token},
		"phrase":       []string{s.Phrase},
		"base64":       []string{base64},
	}
}

func (r PhoneticTranslatorRequest) GetUrl(token string) (string, error) {
	l, ok := availableLangs[r.Language]
	if !ok {
		return "", errors.New("Language not found")
	}

	base, err := url.Parse(fmt.Sprintf(urlQuery, l))
	if err != nil {
		return "", err
	}

	base.RawQuery = r.GetParams(token).Encode()

	return base.String(), nil
}

type PhoneticTranslatorResponse struct {
	Query struct {
		AccessToken                string `json:"access_token"`
		Phrase                     string `json:"phrase"`
		SpellNumbers               int    `json:"spell_numbers"`
		ShowRarePronunciations     int    `json:"show_rare_pronunciations"`
		SplitIntoSyllables         int    `json:"split_into_syllables"`
		AddAspirationSymbol        int    `json:"add_aspiration_symbol"`
		CotCaughtMerger            int    `json:"cot_caught_merger"`
		PinPenMerger               int    `json:"pin_pen_merger"`
		RReplacement               int    `json:"r_replacement"`
		ErReplacement              int    `json:"er_replacement"`
		NarrowTranscription        int    `json:"narrow_transcription"`
		OnlyIForEsEdEndings        int    `json:"only_i_for_es_ed_endings"`
		EnglishPhoneticsAlgorithm  string `json:"english_phonetics_algorithm"`
		ConvertToEnglish           string `json:"Convert_to_english"`
		ElongationSymbolAfterIAndU string `json:"elongation_symbol_after_i_and_u"`
		RColoredVowels             string `json:"r_colored_vowels"`
		AustralianVowels           string `json:"australian_vowels"`
	} `json:"query"`
	PhoneticTranscription []struct {
		Type              string   `json:"type"`
		PunctuationBefore string   `json:"punctuation_before"`
		PunctuationAfter  string   `json:"punctuation_after"`
		Word              string   `json:"word"`
		Transcriptions    []string `json:"transcriptions"`
	} `json:"phonetic_transcription"`
}
