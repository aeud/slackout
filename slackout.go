package slackout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	Username string
	Icon     string
}

var W = SlackOutput{
	Endpoint: os.Getenv("SLACK_ENDPOINT"),
	Username: os.Getenv("SLACK_USERNAME"),
	Icon:     os.Getenv("SLACK_ICON"),
}

func init() {
	log.SetOutput(W)
}

func (s SlackOutput) Write(p []byte) (n int, err error) {
	if W.Endpoint == "" {
		return os.Stdout.Write(p)
	}
	text := fmt.Sprintf("```%s```", string(p))

	if s := os.Getenv("HOSTNAME"); s != "" {
		text = fmt.Sprintf("> po/%s\n%s", s, text)
	}
	if s := os.Getenv("SLACK_JOB_NAME"); s != "" {
		text = fmt.Sprintf("> jobs/%s\n%s", s, text)
	}
	bs, err := json.Marshal(Payload{
		Username: s.Username,
		Text:     text,
		Icon:     s.Icon,
	})
	if err != nil {
		return 0, err
	}
	http.Post(s.Endpoint, "application/json", bytes.NewReader(bs))
	return len(p), nil
}
