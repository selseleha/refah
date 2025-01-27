package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) FundsState(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {

	case resource.BtnStaffFutureFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeFundsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnSupportFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeFundsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnSpecialSavingsFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeFundsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnParsFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeFundsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      MainMenuStateId,
			TextMessage:  resource.MainMenuMessage,
			ReplyButtons: resource.GenerateMainMenuButtons(),
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
