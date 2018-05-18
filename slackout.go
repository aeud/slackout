package slackout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Payload struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Icon     string `json:"icon_emoji"`
}

type SlackOutput struct {
	Endpoint string
	Icon     string
	Username string
}

var W = SlackOutput{
	Endpoint: os.Getenv("SLACK_ENDPOINT"),
}

func (s SlackOutput) Write(p []byte) (n int, err error) {
	bs, err := json.Marshal(Payload{
		Username: s.Username,
		Text:     string(p),
		Icon:     s.Icon,
	})
	if err != nil {
		return 0, err
	}
	http.Post(s.Endpoint, "application/json", bytes.NewReader(bs))
	return 0, nil
}
