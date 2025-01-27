package bot_states

import (
	"gopkg.in/telebot.v3"
	"log"
	"refah/chain"
	"time"
)

func (rs *RefahStates) goToState(node *chain.Node, message *telebot.Message, params *StateParams) *chain.Node {
	newNode, _ := node.GetFlow().Search(params.StateId)

	defer func() {
		if params.Invoice != nil {

		} else if params.Photo != nil {
			var err error
			if params.InlineButtons != nil {
				err = rs.Sender.SendPhoto(node.GetFlow().GetBot(), message.Sender, params.Photo, &telebot.SendOptions{ReplyMarkup: &telebot.ReplyMarkup{InlineKeyboard: params.InlineButtons}})
			} else if params.ReplyButtons != nil {
				err = rs.Sender.SendPhoto(node.GetFlow().GetBot(), message.Sender, params.Photo, &telebot.SendOptions{ReplyMarkup: &telebot.ReplyMarkup{ReplyKeyboard: params.ReplyButtons}})
			} else {
				err = rs.Sender.SendPhoto(node.GetFlow().GetBot(), message.Sender, params.Photo, nil)
			}
			if err != nil {
				log.Fatal("Error sending message: ", err)
			}

		} else {
			var err error
			if params.InlineButtons != nil {
				err = rs.Sender.SendText(node.GetFlow().GetBot(), message.Sender, params.TextMessage, &telebot.SendOptions{ReplyMarkup: &telebot.ReplyMarkup{InlineKeyboard: params.InlineButtons}})
			} else {
				err = rs.Sender.SendText(node.GetFlow().GetBot(), message.Sender, params.TextMessage, &telebot.SendOptions{ReplyMarkup: &telebot.ReplyMarkup{ReplyKeyboard: params.ReplyButtons}})
			}
			if err != nil {
				log.Fatal("Error sending message: ", err)
			}
		}
	}()

	extraInfos := make(map[string]interface{})
	extraInfos["current_state"] = node.GetId()
	extraInfos["next_state"] = params.StateId

	return newNode
}

func (rs *RefahStates) goToStateFromInlineBtn(flow *chain.Chain, receiver *telebot.User, params *StateParams, startTime time.Time, forceReply ...bool) {
	newNode, _ := flow.Search(params.StateId)

	currentStateId := "go_to_state_from_inline_btn"
	if currentNode, ok := flow.GetPosition(receiver); ok {
		currentStateId = currentNode.GetId()
	}

	err := rs.Sender.SendText(flow.GetBot(), receiver, params.TextMessage, &telebot.SendOptions{ReplyMarkup: &telebot.ReplyMarkup{ReplyKeyboard: params.ReplyButtons}})
	if err != nil {
		log.Fatal("Error sending message: ", err)
	}

	extraInfos := make(map[string]interface{})
	extraInfos["next_state"] = params.StateId
	extraInfos["current_state"] = currentStateId

	flow.SetPosition(receiver, newNode)
}
