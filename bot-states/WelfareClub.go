package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) WelfareClubState(node *chain.Node, message *telebot.Message) *chain.Node {
	switch message.Text {

	case resource.BtnHowToReceiveCredit:
		return rs.goToState(node, message, &StateParams{
			StateId:      WelfareClubStateId,
			TextMessage:  resource.ReceiveCreditMessage,
			ReplyButtons: resource.GenerateWelfareClubButtons(),
		})

	case resource.BtnCreditBalance:
		return rs.goToState(node, message, &StateParams{
			StateId:      WelfareClubStateId,
			TextMessage:  resource.BalanceDetailsMessage,
			ReplyButtons: resource.GenerateWelfareClubButtons(),
		})

	case resource.BtnContractCenters:
		return rs.goToState(node, message, &StateParams{
			StateId:      ContractCenterId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateContractCenterButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      MainMenuStateId,
			TextMessage:  resource.MainMenuMessage,
			ReplyButtons: resource.GenerateMainMenuButtons(),
		})

	}
	return rs.wrongStateAndBackToMainMenu(node, message, resource.GenerateMainMenuButtons(), nil)

}
