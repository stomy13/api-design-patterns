package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	marshalJSON()
	unmarshalJSON()
}

func marshalJSON() {
	fmt.Println("### Marshal JSON")
	chatRoom := ChatRoom{
		Messages: MessageList{
			TextMessage{
				ID:   "1",
				Type: "text",
				Text: "Hello, World!",
			},
			ImageMessage{
				ID:       "2",
				Type:     "image",
				ImageURI: "https://example.com/image.jpg",
			},
			VideoMessage{
				ID:       "3",
				Type:     "video",
				VideoURI: "https://example.com/video",
			},
		},
	}
	b, err := json.Marshal(chatRoom)
	if err != nil {
		panic(err)
	}
	jsonStr := string(b)
	fmt.Println(jsonStr)
}

func unmarshalJSON() {
	fmt.Println("### Unmarshal JSON")
	jsonStr := `{"Messages":[{"ID":"4","Type":"text","Text":"Hello, World!"},{"ID":"5","Type":"image","ImageURI":"https://example.com/image.jpg"},{"ID":"6","Type":"video","VideoURI":"https://example.com/video"}]}`
	var chatRoom ChatRoom
	err := json.Unmarshal([]byte(jsonStr), &chatRoom)
	if err != nil {
		panic(err)
	}
	// Polymorphism in Go is not efficient for data structure only
	// This prints slices of map[string]interface{}
	fmt.Printf("%+v\n", chatRoom)
}

type ChatRoom struct {
	Messages MessageList
}

// MessageList is a slice of Message interface
type MessageList []interface{}

type Message interface {
	GetID() MessageID
	GetType() string
}

type MessageID string

type TextMessage struct {
	ID   MessageID
	Type string
	Text string
}

var _ Message = TextMessage{}

func (m TextMessage) GetID() MessageID {
	return m.ID
}

func (m TextMessage) GetType() string {
	return m.Type
}

type ImageMessage struct {
	ID       MessageID
	Type     string
	ImageURI string
}

var _ Message = ImageMessage{}

func (m ImageMessage) GetID() MessageID {
	return m.ID
}

func (m ImageMessage) GetType() string {
	return m.Type
}

type VideoMessage struct {
	ID       MessageID
	Type     string
	VideoURI string
}

var _ Message = VideoMessage{}

func (m VideoMessage) GetID() MessageID {
	return m.ID
}

func (m VideoMessage) GetType() string {
	return m.Type
}
