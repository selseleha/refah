package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) SubmitSurveyState(node *chain.Node, message *telebot.Message) *chain.Node {

	return rs.goToState(node, message, &StateParams{
		StateId:      MainMenuStateId,
		TextMessage:  "امتیاز شما ثبت شد",
		ReplyButtons: resource.GenerateMainMenuButtons(),
	})
}
