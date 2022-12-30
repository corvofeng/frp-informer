package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/blinkbean/dingtalk"
)

func TestXXX(t *testing.T) {
	dingToken := os.Getenv("DINGTALK_TOKEN")
	dingSecret := os.Getenv("DINGTALK_SECRET")
	fmt.Println("get ding token", dingToken, dingSecret)
	cli := dingtalk.InitDingTalkWithSecret(
		dingToken, dingSecret,
	)
	// cli.SendTextMessage("content")
	msg := []string{
		"### Link test",
		"---",
		"- <font color=#00ff00 size=6>红色文字</font>",
		"- content2",
	}
	cli.SendMarkDownMessageBySlice("Markdown title", msg, dingtalk.WithAtAll())

}
