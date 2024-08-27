package requests

import (
	"Zametki/configs"
	"Zametki/models"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func spellChecker(text string) ([]models.SpellError, error) {
	data := url.Values{}
	data.Set("text", text)
	data.Set("lang", "ru")

	resp, err := http.Post(configs.GetEnv("SPELLER_URL"), "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return []models.SpellError{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.SpellError{}, err
	}

	var errors []models.SpellError
	err = json.Unmarshal(body, &errors)
	if err != nil {
		return []models.SpellError{}, err
	}

	return errors, nil
}

func formatText(text string, errors []models.SpellError) string {
	var result strings.Builder
	currentPos := 0

	runes := []rune(text)

	for _, err := range errors {
		result.WriteString(string(runes[currentPos:err.Pos]))

		if len(err.S) > 0 {
			result.WriteString(err.S[0])
		} else {
			result.WriteString(string(runes[err.Pos : err.Pos+err.Len]))
		}

		currentPos = err.Pos + err.Len
	}

	result.WriteString(string(runes[currentPos:]))

	return result.String()
}
