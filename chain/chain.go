package chain

/*
	Chain flow is list of event listeners organized by type for Telegram
	Author: Daniil Furmanov
	License: MIT
*/

import (
	"errors"
	"gopkg.in/telebot.v3"
	"sync"
)

/*
A flow is chain or double-linked list of events organized by type
*/
type Chain struct {
	id                     string
	root                   *Node
	bot                    *telebot.Bot
	defaultLocale          string
	positions              map[string]*Node
	defaultMessageHandler  MessageEndpoint
	defaultCallbackHandler CallbackEndpoint
	mx                     sync.RWMutex
}

var ErrChainIsEmpty = errors.New("chain has zero handlers")

/*
Creates a new chain flow
*/
func NewChainFlow(id string, bot *telebot.Bot) (*Chain, error) {
	f := &Chain{
		id:                     id,
		bot:                    bot,
		positions:              make(map[string]*Node),
		defaultMessageHandler:  nil,
		defaultCallbackHandler: nil,
		mx:                     sync.RWMutex{},
	}
	f.root = &Node{id: id + "_root", flow: f, messageEndpoint: nil, callbackEndpoint: nil, prev: nil, next: nil}
	return f, nil
}

/*
Get chain's unique identificator
*/
func (c *Chain) GetId() string {
	return c.id
}

/*
Get attached Telegram bot
*/
func (c *Chain) GetBot() *telebot.Bot {
	return c.bot
}

/*
Get the root node
*/
func (c *Chain) GetRoot() *Node {
	return c.root
}

/*
Gets the user position in the flow
*/
func (c *Chain) GetPosition(of telebot.Recipient) (*Node, bool) {
	c.mx.RLock()
	node, ok := c.positions[of.Recipient()]
	c.mx.RUnlock()
	return node, ok
}

/*
Sets the user current position in the flow
*/
func (c *Chain) SetPosition(of telebot.Recipient, node *Node) {
	c.mx.Lock()
	c.positions[of.Recipient()] = node
	c.mx.Unlock()
}

/*
Deletes the user current position in the flow
*/
func (c *Chain) DeletePosition(of telebot.Recipient) {
	c.mx.Lock()
	delete(c.positions, of.Recipient())
	c.mx.Unlock()
}

/*
Search for a node with ID
*/
func (c *Chain) Search(nodeId string) (*Node, bool) {
	return c.root.SearchDown(nodeId)
}

/*
Get the root node
*/
func (c *Chain) SetDefaultMessageHandler(endpoint MessageEndpoint) *Chain {
	c.defaultMessageHandler = endpoint
	return c
}

func (c *Chain) SetDefaultCallbackHandler(endpoint CallbackEndpoint) *Chain {
	c.defaultCallbackHandler = endpoint
	return c
}

/*
Executes the chain for the user by putting him on a first stage of the chain
*/
func (c *Chain) Start(to telebot.Recipient, text string, options ...interface{}) (err error) {
	if c.root.next == nil {
		return ErrChainIsEmpty
	}
	if options != nil && len(options) > 0 {
		// a workaround for nil options
		// otherwise the message will not be sent
		_, err = c.GetBot().Send(to, text, options...)
	} else {
		_, err = c.GetBot().Send(to, text)
	}
	if err == nil {
		c.SetPosition(to, c.root.next)
	}
	return
}

func (c *Chain) StartWithPhoto(to telebot.Recipient, photo *telebot.Photo, options ...interface{}) (err error) {
	if c.root.next == nil {
		return ErrChainIsEmpty
	}
	if options != nil && len(options) > 0 {
		// a workaround for nil options
		// otherwise the message will not be sent
		_, err = c.GetBot().Send(to, photo, options...)
	} else {
		_, err = c.GetBot().Send(to, photo)
	}
	if err == nil {
		c.SetPosition(to, c.root.next)
	}
	return
}

/*
Process with the next flow iteration
Returns true only if the iteration was successful
*/
func (c *Chain) Process(m *telebot.Message, callback *telebot.Callback) bool {
	if m != nil {
		sender := m.Sender
		node, ok := c.GetPosition(sender)
		if !ok {
			// the flow hasn't started for the user
			return false
		}
		if node == nil {
			c.DeletePosition(sender)
			return false
		}
		if !node.CheckEvent(m) || node.messageEndpoint == nil {
			// input is invalid for the particular node
			if c.defaultMessageHandler != nil {
				next := c.defaultMessageHandler(node, m)
				if next != node {
					c.SetPosition(sender, next)
				}
				return true
			}
			return false
		}
		next := node.messageEndpoint(node, m)
		if next != node {
			c.SetPosition(sender, next)
		}
		return true
	}

	if callback != nil {
		sender := callback.Sender
		node, ok := c.GetPosition(sender)
		if !ok {
			// the flow hasn't started for the user
			return false
		}
		if node == nil {
			c.DeletePosition(sender)
			return false
		}
		if node.callbackEndpoint == nil {
			// input is invalid for the particular node
			if c.defaultCallbackHandler != nil {
				next := c.defaultCallbackHandler(node, callback)
				if next != node {
					c.SetPosition(sender, next)
				}
				return true
			}
			return false
		}
		next := node.callbackEndpoint(node, callback)
		if next != node {
			c.SetPosition(sender, next)
		}
		return true
	}
	return true
}
