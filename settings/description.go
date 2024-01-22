package settings

import (
	"encoding/json"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

type description struct {
	Description string `json:"description"`
}

func GetBotDescription(bot tele.Bot) (string, error) {
	data, err := bot.Raw("getMyDescription", nil)
	if err != nil {
		return "", fmt.Errorf("getMyDescription: %w", err)
	}
	var resp struct {
		Ok     bool        `json:"ok"`
		Result description `json:"result"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", fmt.Errorf("json unmarshal: %w", err)
	}

	return resp.Result.Description, nil
}
