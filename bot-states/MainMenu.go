package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) MainMenuState(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {
	case resource.BtnWelfareClub:
		return rs.goToState(node, message, &StateParams{
			StateId:      WelfareClubStateId,
			TextMessage:  resource.SelectMessage,
			ReplyButtons: resource.GenerateWelfareClubButtons(),
		})

	case resource.BtnFunds:
		return rs.goToState(node, message, &StateParams{
			StateId:      FundsStateId,
			TextMessage:  resource.SelectMessage,
			ReplyButtons: resource.GenerateFundsButtons(),
		})

	case resource.BtnSurvey:
		return rs.goToState(node, message, &StateParams{
			StateId:      SurveyStateId,
			TextMessage:  resource.SelectMessage3,
			ReplyButtons: resource.GenerateSurveyButtons(),
		})
	}

	return rs.wrongStateAndBackToMainMenu(node, message, resource.GenerateMainMenuButtons(), nil)
}
