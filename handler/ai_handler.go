package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/eatmoreapple/openwechat"
)

type QWenResp struct {
	Output    Output `json:"output"`
	Usage     Usage  `json:"usage"`
	RequestID string `json:"request_id"`
}

type Output struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	TotalTokens  int64 `json:"total_tokens"`
	OutputTokens int64 `json:"output_tokens"`
	InputTokens  int64 `json:"input_tokens"`
}

func AiMatchFunc(m *openwechat.Message) bool {
	return m.IsAt() && strings.Contains(m.Content, "@robot") && !strings.Contains(m.Content, "周文喆")
}

func AiHandler(ctx *openwechat.MessageContext) {
	userContent := strings.ReplaceAll(ctx.Content, "@robot", "")
	req, err := http.NewRequest("POST",
		"https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
		strings.NewReader(fmt.Sprintf(`{
			"model": "qwen-turbo",
			"input": {
				"messages": [
					{
						"role": "system",
						"content": "You are a helpful assistant."
					},
					{
						"role": "user",
						"content": "%s"
					}
				]
			},
			"parameters": {
				"result_format": "message"
			}
		}`, userContent)))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer sk-7509c5aeefd1455db8a15bdcdb49d79b")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "acw_tc=03445b3a-ebd5-9d72-a3c8-3e4c334c1a87a605cddf19d5be46f524361c193025d3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var qwenResp QWenResp
	err = json.Unmarshal(body, &qwenResp)
	if err != nil {
		log.Println(err)
		return
	}

	choiceList := qwenResp.Output.Choices

	if len(choiceList) == 0 {
		log.Println("没有匹配到内容")
		return
	}
	ctx.ReplyText(choiceList[0].Message.Content)
}
