package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) DescribeCompany(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {

	case resource.BtnIntroductionCompany:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeCompanyStateId,
			ReplyButtons: resource.GenerateCompaniesDescribeButtons(),
			Photo: &telebot.Photo{
				File:    resource.SafarMarketImage,
				Caption: resource.SafarMarketMessage,
			},
		})

	case resource.BtnConsumptionReport:
		return rs.goToState(node, message, &StateParams{
			StateId:      DescribeCompanyStateId,
			TextMessage:  resource.SafarMarketReportMessage,
			ReplyButtons: resource.GenerateCompaniesDescribeButtons(),
		})

	case resource.BtnRefund:
		return rs.goToState(node, message, &StateParams{
			StateId:      CompaniesStateId,
			TextMessage:  resource.SelectMessage,
			ReplyButtons: resource.GenerateCompaniesButtons(),
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
