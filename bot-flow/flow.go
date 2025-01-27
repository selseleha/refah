package bot_flow

import (
	"gopkg.in/telebot.v3"
	"refah/bot-states"
	"refah/chain"
)

type StateData struct {
	StateId          string
	MessageEndpoint  chain.MessageEndpoint
	CallbackEndpoint chain.CallbackEndpoint
	ExpectedEvent    []string
}

var StateDataList []StateData

func CreateFlow(bot *telebot.Bot, settings *bot_states.RefahSettings) *chain.Chain {

	refahStates := bot_states.NewRefahStatesImpl(settings)

	StateDataList = []StateData{
		{bot_states.MainMenuStateId, refahStates.MainMenuState, nil, []string{telebot.OnText}},
		{bot_states.WelfareClubStateId, refahStates.WelfareClubState, nil, []string{telebot.OnText}},
		{bot_states.ContractCenterId, refahStates.ContractCentersState, nil, []string{telebot.OnText}},
		{bot_states.SurveyStateId, refahStates.SurveyState, nil, []string{telebot.OnText}},
		{bot_states.SubmitSurveyStateId, refahStates.SubmitSurveyState, nil, []string{telebot.OnText}},
		{bot_states.HotelsStateId, refahStates.HotelsState, nil, []string{telebot.OnText}},
		{bot_states.FundsStateId, refahStates.FundsState, nil, []string{telebot.OnText}},
		{bot_states.DescribeFundsStateId, refahStates.DescribeFunds, nil, []string{telebot.OnText}},
		{bot_states.CompaniesStateId, refahStates.CompaniesState, nil, []string{telebot.OnText}},
		{bot_states.DescribeCompanyStateId, refahStates.DescribeCompany, nil, []string{telebot.OnText}},
	}

	flow, _ := chain.NewChainFlow("refah", bot)
	flow.SetDefaultCallbackHandler(nil)
	lastNode := flow.GetRoot()

	for _, state := range StateDataList {
		lastNode = lastNode.Then(state.StateId, state.MessageEndpoint, state.CallbackEndpoint, state.ExpectedEvent...)
	}

	refahStates.InitHandlers(flow)
	return flow
}
