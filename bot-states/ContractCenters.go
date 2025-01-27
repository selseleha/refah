package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) ContractCentersState(node *chain.Node, message *telebot.Message) *chain.Node {
	switch message.Text {

	case resource.BtnHotels:
		return rs.goToState(node, message, &StateParams{
			StateId:      HotelsStateId,
			TextMessage:  resource.SelectHotelMessage,
			ReplyButtons: resource.GenerateHotelsButtons(),
		})

	case resource.BtnCompanies:
		return rs.goToState(node, message, &StateParams{
			StateId:      CompaniesStateId,
			TextMessage:  resource.SelectCompanyMessage,
			ReplyButtons: resource.GenerateCompaniesButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      WelfareClubStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateWelfareClubButtons(),
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
