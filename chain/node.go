package chain

import (
	"gopkg.in/telebot.v3"
)

/*
Callback function declaration that includes a "node" that a call was made from
*/
type MessageEndpoint func(e *Node, c *telebot.Message) *Node

type CallbackEndpoint func(e *Node, c *telebot.Callback) *Node

/*
Node is an element in a double-linked list
*/
type Node struct {
	id               string
	flow             *Chain
	messageEndpoint  MessageEndpoint
	callbackEndpoint CallbackEndpoint
	prev             *Node
	next             *Node
	event            []string
}

/*
Creates a following element in the list
*/
func (e *Node) Then(id string, endpoint MessageEndpoint, callbackEndpoint CallbackEndpoint, expectedEvents ...string) *Node {
	newNode := &Node{
		id:               id,
		flow:             e.flow,
		messageEndpoint:  endpoint,
		callbackEndpoint: callbackEndpoint,
		prev:             e,
		next:             nil,
		event:            expectedEvents,
	}
	e.next = newNode
	return newNode
}

/*
Get related flow
*/
func (e *Node) GetFlow() *Chain {
	return e.flow
}

/*
Get node's identificator
*/
func (e *Node) GetId() string {
	return e.id
}

/*
Get node's callback endpoint
*/
func (e *Node) GetEndpoint() MessageEndpoint {
	return e.messageEndpoint
}

func (e *Node) GetCallbackEndpoint() CallbackEndpoint {
	return e.callbackEndpoint
}

/*
Get the previous node in the list
*/
func (e *Node) Previous() *Node {
	return e.prev
}

/*
Get the next node in the list
*/
func (e *Node) Next() *Node {
	return e.next
}

/*
Tries to find a node with ID down the list
*/
func (e *Node) SearchDown(nodeId string) (*Node, bool) {
	temp := e
	for {
		temp = temp.next
		if temp == nil {
			break
		}
		if temp.id == nodeId {
			return temp, true
		}
	}
	return nil, false
}

/*
Tries to find a node with ID up the list
*/
func (e *Node) SearchUp(nodeId string) (*Node, bool) {
	temp := e
	for {
		temp = temp.prev
		if temp == nil {
			break
		}
		if temp.id == nodeId {
			return temp, true
		}
	}
	return nil, false
}

/*
Checks if the message type is matching the node type
*/
func (e *Node) CheckEvent(m *telebot.Message) bool {
	for _, event := range e.event {
		switch event {
		case telebot.OnText:
			if len(m.Text) > 0 {
				return true
			}
		case telebot.OnPhoto:
			if m.Photo != nil {
				return true
			}
		case telebot.OnLocation:
			if m.Location != nil {
				return true
			}
		case telebot.OnContact:
			if m.Contact != nil {
				return true
			}
		case telebot.OnAudio:
			if m.Audio != nil {
				return true
			}
		case telebot.OnVideoNote:
			if m.VideoNote != nil {
				return true
			}
		case telebot.OnVideo:
			if m.Video != nil {
				return true
			}
		case telebot.OnVoice:
			if m.Voice != nil {
				return true
			}
		case telebot.OnDocument:
			if m.Document != nil {
				return true
			}
		case telebot.OnSticker:
			if m.Sticker != nil {
				return true
			}
		}
	}

	return false
}
