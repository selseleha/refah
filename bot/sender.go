package bot

import (
	"errors"
	"gopkg.in/telebot.v3"
	"log"
)

type Sender interface {
	SendText(bot *telebot.Bot, receiver *telebot.User, message string, option *telebot.SendOptions) error
	SendPhoto(bot *telebot.Bot, receiver *telebot.User, photo *telebot.Photo, options *telebot.SendOptions) error
	EditTextMessage(bot *telebot.Bot, msg telebot.Editable, text string, option *telebot.ReplyMarkup)
}

type MessageSender struct {
}

func NewMessageSender() *MessageSender {
	return &MessageSender{}
}

func (s *MessageSender) SendText(bot *telebot.Bot, receiver *telebot.User, message string, options *telebot.SendOptions) error {
	var err error
	var res *telebot.Message
	if options == nil {
		res, err = bot.Send(receiver, message)
	} else {
		res, err = bot.Send(receiver, message, options)
	}
	if err != nil {
		log.Fatalf("Failed to send message : %s  receiverId: %s", message, receiver.ID)
		return err
	}

	if res == nil {
		log.Fatalf("result is nil")

		return errors.New("cant send message")
	}
	return nil
}

func (s *MessageSender) SendPhoto(bot *telebot.Bot, receiver *telebot.User, photo *telebot.Photo, options *telebot.SendOptions) error {
	var err error
	var res *telebot.Message
	if options == nil {
		res, err = bot.Send(receiver, photo)
	} else {
		res, err = bot.Send(receiver, photo, options)
	}
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("cant send photo by tapi.bale.ai")
	}
	return nil
}

func (s *MessageSender) EditTextMessage(bot *telebot.Bot, message telebot.Editable, text string, option *telebot.ReplyMarkup) {
	message.MessageSig()
	if res, err := bot.Edit(message, text, option); err != nil {
		log.Fatalf("error in editing message")
		return
	} else if res == nil {

		log.Fatalf("error in editing message")
	}
}
