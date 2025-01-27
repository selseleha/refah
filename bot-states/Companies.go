package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) CompaniesState(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {

	case resource.BtnSafarMarket:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeCompanyStateId,
			TextMessage:  resource.SelectMessage4,
			ReplyButtons: resource.GenerateCompaniesDescribeButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      ContractCenterId,
			TextMessage:  resource.SelectMessage,
			ReplyButtons: resource.GenerateContractCenterButtons(),
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
