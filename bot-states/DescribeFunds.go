package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) DescribeFunds(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {

	case resource.BtnIntroductionFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeFundsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnSiteFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      HotelsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnLastFileFund:
		return rs.goToState(node, message, &StateParams{
			StateId:      HotelsStateId,
			TextMessage:  resource.SelectMessage2,
			ReplyButtons: resource.GenerateFundDescribeButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      FundsStateId,
			TextMessage:  resource.SelectMessage,
			ReplyButtons: resource.GenerateFundsButtons(),
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
