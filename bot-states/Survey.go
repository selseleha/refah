package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) SurveyState(node *chain.Node, message *telebot.Message) *chain.Node {
	switch message.Text {

	case resource.BtnMedicalCenter:
		return rs.goToState(node, message, &StateParams{
			StateId:       SubmitSurveyStateId,
			TextMessage:   resource.SatisfactionMessage,
			InlineButtons: resource.GenerateSurveyInlineButtons(),
		})

	case resource.BtnJahanHotel:
		return rs.goToState(node, message, &StateParams{
			StateId:       MainMenuStateId,
			TextMessage:   resource.SatisfactionMessage,
			InlineButtons: resource.GenerateSurveyInlineButtons(),
		})

	case resource.BtnAtiyehHotel:
		return rs.goToState(node, message, &StateParams{
			StateId:       WelfareClubStateId,
			TextMessage:   resource.SatisfactionMessage,
			InlineButtons: resource.GenerateSurveyInlineButtons(),
		})

	case resource.BtnBackToMainMenu:
		return rs.goToState(node, message, &StateParams{
			StateId:      MainMenuStateId,
			TextMessage:  resource.MainMenuMessage,
			ReplyButtons: resource.GenerateMainMenuButtons(),
		})
	}
	return rs.wrongStateAndBackToMainMenu(node, message, resource.GenerateMainMenuButtons(), nil)

}
