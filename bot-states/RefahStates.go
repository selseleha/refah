package bot_states

import (
	"gopkg.in/telebot.v3"
	"log"
	"refah/bot"
	resource "refah/bot-resources"
	"refah/chain"
	"time"
)

type RefahSettings struct {
	Sender bot.Sender
}

type RefahStates struct {
	Sender bot.Sender
}

func NewRefahStatesImpl(settings *RefahSettings) *RefahStates {
	return &RefahStates{
		Sender: settings.Sender,
	}
}

type StateParams struct {
	StateId       string
	TextMessage   string
	ReplyButtons  [][]telebot.ReplyButton
	Invoice       *telebot.Invoice
	InlineButtons [][]telebot.InlineButton
	Photo         *telebot.Photo
}

const (
	MainMenuStateId           = "main_menu"
	WelfareClubStateId        = "welfare_club"
	FundsStateId              = "funds"
	DescribeFundsStateId      = "describe_funds"
	HealthAndTreatmentStateId = "health_and_treatment"
	CompensationStateId       = "compensation"
	SportsStateId             = "sports"
	GeneralNewsStateId        = "general_news"
	SurveyStateId             = "survey"
	SubmitSurveyStateId       = "submit_survey"
	ContractCenterId          = "contract_center"
	HotelsStateId             = "hotels"
	CompaniesStateId          = "companies"
	DescribeCompanyStateId    = "describe_companies"

	BackToMainMenuStateId = "back_to_main_menu_state"
)

// var ScoreInlineBtn = telebot.InlineButton{Unique: "handle_score"}
var ScoreInlineBtn = resource.GenerateSurveyInlineButtons()

func (rs *RefahStates) InitHandlers(flow *chain.Chain) {
	b := flow.GetBot()

	b.Handle("/start", func(c telebot.Context) error {
		m := c.Message()

		rs.startBot(m, flow, true)
		return nil
	})

	supportedFormat := []string{telebot.OnText, telebot.OnContact, telebot.OnLocation, telebot.OnPhoto, telebot.OnVideo, telebot.OnSticker, telebot.OnPayment, telebot.OnVoice}

	for _, format := range supportedFormat {
		currentFormat := format
		b.Handle(currentFormat, func(c telebot.Context) error {
			m := c.Message()
			_, hasPosition := flow.GetPosition(m.Sender)

			if hasPosition {
				flow.Process(m, nil)
			} else {
				rs.startBot(m, flow, m.Text == "/start")
			}

			return nil
		})
	}

	b.Handle(telebot.OnCallback, func(c telebot.Context) error {
		//
		//switch c.Callback().Data {
		//case "score_verylow":
		//	return c.Send("امتیاز شما: خیلی کم")
		//case "score_low":
		//	return c.Send("امتیاز شما: کم")
		//case "score_average":
		//	return c.Send("امتیاز شما: متوسط")
		//case "score_good":
		//	return c.Send("امتیاز شما: خوب")
		//case "score_verygood":
		//	return c.Send("امتیاز شما: خیلی خوب")
		//default:
		//	return c.Send("انتخاب نامعتبر")
		//}

		_, hasPosition := flow.GetPosition(c.Sender())

		if hasPosition {
			log.Println("Processing callback in flow...") // Debugging line
			flow.Process(&telebot.Message{
				Sender: c.Sender(),
				Text:   c.Callback().Data,
			}, nil)
		} else {
			log.Println("User has no position in flow!") // Debugging line
		}

		return c.Respond()
	})

	rs.HandleInlineButtons(b, flow)
}

func (rs *RefahStates) stayInStateAndSendEnterCorrectInputMessage(node *chain.Node, message *telebot.Message, buttons [][]telebot.ReplyButton, inlineButtons [][]telebot.InlineButton) *chain.Node {
	return rs.goToState(node, message, &StateParams{
		StateId:       node.GetId(),
		TextMessage:   "اشتباه",
		ReplyButtons:  buttons,
		InlineButtons: inlineButtons,
	})
}

func (rs *RefahStates) wrongStateAndBackToMainMenu(node *chain.Node, message *telebot.Message, buttons [][]telebot.ReplyButton, inlineButtons [][]telebot.InlineButton) *chain.Node {
	return rs.goToState(node, message, &StateParams{
		StateId:      MainMenuStateId,
		TextMessage:  "این قسمت در حال حاضر پیاده سازی نشده است",
		ReplyButtons: resource.GenerateMainMenuButtons(),
	})
}

//func (rs *RefahStates) HandleInlineButtons(bot *telebot.Bot, flow *chain.Chain) {
//	bot.Handle(&ScoreInlineBtn, func(c telebot.Context) error {
//		startTime := time.Now()
//
//		for _, score := range resource.GetScoreButtons() {
//			log.Println(score)
//
//			rs.goToStateFromInlineBtn(flow, c.Sender(), &StateParams{
//				StateId:      SubmitSurveyStateId,
//				TextMessage:  "امتیاز ثبت شد",
//				ReplyButtons: resource.GenerateMainMenuButtons(),
//			}, startTime, false)
//
//			removeScoreButtonsFromInlineMessage(bot, c.Callback(), &rs.Sender)
//		}
//		return nil
//	})
//}

func (rs *RefahStates) HandleInlineButtons(bot *telebot.Bot, flow *chain.Chain) {
	// Handle each score selection
	bot.Handle(&telebot.InlineButton{Unique: "score_verylow"}, func(c telebot.Context) error {
		return rs.handleSurveyScore(c, flow, "خیلی کم")
	})
	bot.Handle(&telebot.InlineButton{Unique: "score_low"}, func(c telebot.Context) error {
		return rs.handleSurveyScore(c, flow, "کم")
	})
	bot.Handle(&telebot.InlineButton{Unique: "score_average"}, func(c telebot.Context) error {
		return rs.handleSurveyScore(c, flow, "متوسط")
	})
	bot.Handle(&telebot.InlineButton{Unique: "score_good"}, func(c telebot.Context) error {
		return rs.handleSurveyScore(c, flow, "خوب")
	})
	bot.Handle(&telebot.InlineButton{Unique: "score_verygood"}, func(c telebot.Context) error {
		return rs.handleSurveyScore(c, flow, "خیلی خوب")
	})
}

func (rs *RefahStates) handleSurveyScore(c telebot.Context, flow *chain.Chain, score string) error {
	startTime := time.Now()
	log.Println("User selected score:", score)

	rs.goToStateFromInlineBtn(flow, c.Sender(), &StateParams{
		StateId:      SubmitSurveyStateId,
		TextMessage:  "امتیاز ثبت شد: " + score,
		ReplyButtons: resource.GenerateMainMenuButtons(),
	}, startTime, false)

	removeScoreButtonsFromInlineMessage(c.Bot(), c.Callback(), &rs.Sender)
	return nil
}
func removeScoreButtonsFromInlineMessage(bot *telebot.Bot, c *telebot.Callback, sender *bot.Sender) {
	(*sender).EditTextMessage(bot, c.Message, removeLastLine(c.Message.Text), nil)
}

func removeLastLine(text string) string {
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == '\n' {
			return text[:i]
		}
	}
	return text
}

func (rs *RefahStates) startBot(m *telebot.Message, flow *chain.Chain, b bool) {

	startMessage := resource.MainMenuMessage
	startReplyMarkup := &telebot.ReplyMarkup{ReplyKeyboard: resource.GenerateMainMenuButtons()}
	if err := flow.Start(m.Sender, startMessage, &telebot.SendOptions{ReplyMarkup: startReplyMarkup}); err != nil {
		log.Println("failed to start the conversation", err)
	}

}
