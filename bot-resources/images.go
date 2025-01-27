package bot_resources

import "gopkg.in/telebot.v3"

const (
	HotelFileId       = "bot-resources/images/hotel.png"
	SafarMarketFileId = "bot-resources/images/safarMarket.png"
)

var (
	HotelImage       = telebot.File{FileLocal: HotelFileId}
	SafarMarketImage = telebot.File{FileLocal: SafarMarketFileId}
)
