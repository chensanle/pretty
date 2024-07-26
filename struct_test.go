package pretty

import (
	"fmt"
	"testing"
)

type DemoStruct struct {
	RequestId string   `json:"requestId"`
	Codes     []int    `json:"code"`
	Reasons   []string `json:"reasons"`
	Data      struct {
		RequestId string `json:"requestId"`
		Choices   []struct {
			Messages []struct {
				Name            string `json:"name"`
				RoleId          string `json:"roleId"`
				Role            string `json:"role"`
				Content         string `json:"content"`
				FinishReason    string `json:"finishReason"`
				ValidMessage    bool   `json:"validMessage"`
				UseMessage      bool   `json:"useMessage"`
				FunctionMessage bool   `json:"functionMessage"`
			} `json:"messages"`
			StopReason string `json:"stopReason"`
		} `json:"choices"`
		Usage struct {
			OutputTokens int `json:"outputTokens"`
			InputTokens  int `json:"inputTokens"`
			UserTokens   int `json:"userTokens"`
			PluginTokens int `json:"pluginTokens"`
		} `json:"usage"`
		Context struct {
			ChatRoomId           int    `json:"chatRoomId"`
			SessionId            string `json:"sessionId"`
			ChatId               string `json:"chatId"`
			AnswerId             string `json:"answerId"`
			MessageId            string `json:"messageId"`
			QueryId              string `json:"queryId"`
			ReplyMessageId       string `json:"replyMessageId"`
			EnableDataInspection bool   `json:"enableDataInspection"`
			IsSave               bool   `json:"isSave"`
			RequestId            string `json:"requestId"`
			ModelRequestId       string `json:"modelRequestId"`
			CharacterPk          int    `json:"characterPk"`
			CharacterName        string `json:"characterName"`
			CharacterId          string `json:"characterId"`
			ModelName            string `json:"modelName"`
			Origin               string `json:"origin"`
			BizSrc               string `json:"bizSrc"`
			Ext                  string `json:"ext"`
			BizUserId            string `json:"bizUserId"`
			ChatLockKey          string `json:"chatLockKey"`
			ResultCount          int    `json:"resultCount"`
			IsGroupChat          bool   `json:"isGroupChat"`
		} `json:"context"`
		Stop bool `json:"stop"`
	} `json:"data"`
	Success bool `json:"success"`
}

func TestValueMap(t *testing.T) {
	demoMap := map[any]any{
		"codes":   []int{1, 2, 3, 4, 5},
		"reasons": []string{"beijing", "nanjing"},
		1:         2,
	}
	fmt.Println(string(Value(demoMap)))
}

func TestValueStruct(t *testing.T) {
	demoStruct := &DemoStruct{
		RequestId: "request_id",
		Codes:     []int{1, 2, 3, 4, 5, 6},
		Reasons:   []string{"beijing", "nanjing"},
	}
	fmt.Print(string(Value(demoStruct)))
}
