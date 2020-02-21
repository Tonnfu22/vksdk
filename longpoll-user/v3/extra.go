package wrapper

import (
	"fmt"
	"reflect"
	"time"
)

type Action struct {
	SourceAct         string
	SourceMid         string
	SourceText        string
	SourceOldText     string
	SourceMessage     string
	SourceChatLocalID string
}

func (result *Action) parse(v map[string]interface{}) {
	if sourceAct, ok := v["source_act"].(string); ok {
		result.SourceAct = sourceAct

		switch sourceAct {
		case "chat_create":
			if sourceText, ok := v["source_text"].(string); ok {
				result.SourceText = sourceText
			}

		case "chat_title_update":
			if sourceText, ok := v["source_text"].(string); ok {
				result.SourceText = sourceText
			}

			if sourceOldText, ok := v["source_old_text"].(string); ok {
				result.SourceOldText = sourceOldText
			}

		case "chat_kick_user", "chat_invite_user":
			if sourceMid, ok := v["source_mid"].(string); ok {
				result.SourceMid = sourceMid
			}

		case "chat_pin_message", "chat_unpin_message":
			if sourceMid, ok := v["source_mid"].(string); ok {
				result.SourceMid = sourceMid
			}

			if sourceMessage, ok := v["source_message"].(string); ok {
				result.SourceMessage = sourceMessage
			}

			if sourceChatLocalID, ok := v["source_chat_local_id"].(string); ok {
				result.SourceChatLocalID = sourceChatLocalID
			}
		}
	}
}

type AdditionalData struct {
	Title     string
	RefSource string
	From      string
	FromAdmin string
	Emoji     string
	Action
}

func (result *AdditionalData) parse(v map[string]interface{}) {
	if title, ok := v["title"].(string); ok {
		result.Title = title
	}

	if refSource, ok := v["ref_source"].(string); ok {
		result.RefSource = refSource
	}

	if from, ok := v["from"].(string); ok {
		result.From = from
	}

	if fromAdmin, ok := v["from_admin"].(string); ok {
		result.FromAdmin = fromAdmin
	}

	if emoji, ok := v["emoji"].(string); ok {
		result.Emoji = emoji
	}

	result.Action.parse(v)
}

type LongPollAttachments map[string]interface{}

// https://vk.com/dev/using_longpoll_3, point 3.1
type ExtraFields struct {
	PeerID         int
	Timestamp      time.Time
	Text           string
	AdditionalData AdditionalData
	Attachments    LongPollAttachments
}

func (result *ExtraFields) parseExtraFields(i []interface{}) error {
	length := len(i)

	if length > 3 {
		if v, ok := i[3].(float64); ok {
			result.PeerID = int(v)
		}
	}

	if length > 4 {
		if v, ok := i[4].(float64); ok {
			result.Timestamp = time.Unix(int64(v), 0)
		}
	}

	if length > 5 {
		if v, ok := i[5].(string); ok {
			result.Text = v
		}
	}

	if length > 6 {
		if v, ok := i[6].(map[string]interface{}); ok {
			result.AdditionalData.parse(v)
		}
	}

	if length > 7 {
		if v, ok := i[7].(map[string]interface{}); ok {
			result.Attachments = v
		}
	}

	return nil
}

func interfaceToStringIntMap(m interface{}) (map[string]int, error) {
	reflectedMap := reflect.ValueOf(m)
	if reflectedMap.Kind() != reflect.Map {
		return nil, fmt.Errorf("expected a slice, got %T", m)
	}

	result := make(map[string]int, reflectedMap.Len())

	for _, key := range reflectedMap.MapKeys() {
		v, ok := reflectedMap.MapIndex(key).Interface().(float64)
		if !ok {
			return nil, fmt.Errorf("cast failed, value type: %T", reflectedMap.MapIndex(key).Interface())
		}

		result[key.String()] = int(v)
	}

	return result, nil
}

func interfaceToIDSlice(slice interface{}) ([]int, error) {
	reflectedSlice := reflect.ValueOf(slice)
	if reflectedSlice.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %T", slice)
	}

	result := make([]int, reflectedSlice.Len())

	for i := 0; i < reflectedSlice.Len(); i++ {
		v, ok := reflectedSlice.Index(i).Interface().(float64)
		if !ok {
			return nil, fmt.Errorf("cast failed, value type: %T", reflectedSlice.Index(i).Interface())
		}

		result[i] = int(v)
	}

	return result, nil
}
