package bot_states

import (
	"gopkg.in/telebot.v3"
	resource "refah/bot-resources"
	"refah/chain"
)

func (rs *RefahStates) HotelsState(node *chain.Node, message *telebot.Message) *chain.Node {

	switch message.Text {

	case resource.BtnParsHotel:
		return rs.goToState(node, message, &StateParams{
			StateId:      HotelsStateId,
			ReplyButtons: resource.GenerateHotelsButtons(),
			Photo: &telebot.Photo{
				File:    resource.HotelImage,
				Caption: resource.ParsHotelMessage,
			},
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
