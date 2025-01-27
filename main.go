package main

import "C"
import (
	"encoding/json"
	"golang.org/x/net/proxy"
	"gopkg.in/telebot.v3"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	b "refah/bot"
	bf "refah/bot-flow"
	bs "refah/bot-states"

	"time"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Dump request for debugging
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Printf("[API REQUEST] %s\n", reqDump)
	}

	// Execute the request
	resp, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Dump response for debugging
	respDump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Printf("[API RESPONSE] %s\n", respDump)
	}

	return resp, err
}

type Configuration struct {
	BotToken string `json:"botToken"`
	BaseUrl  string `json:"baseUrl"`
}

func loadConfig(filename string) (*Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {

	config, err := loadConfig("conf.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	telegram := false
	var pref telebot.Settings

	if telegram {
		proxyAddr := "192.168.10.10:1080" // Replace with your proxy's address and port
		dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			log.Fatalf("Failed to set up SOCKS5 proxy: %v", err)
		}

		httpTransport := &http.Transport{
			Dial: dialer.Dial,
		}

		client := &http.Client{
			Transport: httpTransport,
		}

		//botToken := "8127610510:AAFzxXcaeF9IcRcGZjBCxFcISOxw6DiP5ek"

		pref = telebot.Settings{
			Token:  config.BotToken,
			Client: client,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}
	} else {
		baseURL := config.BaseUrl

		//botToken := "2000000001:UUg9lhQx1i6rxyLxpPzJyXWyFJfYyzpUtA"
		//botToken := "2000000007:3Lsv7gL7qY50efVIOBzjLz9D3dHU9UVDKA"

		client := &http.Client{
			Transport: LoggingRoundTripper{Proxied: http.DefaultTransport},
			Timeout:   10 * time.Second,
		}

		pref = telebot.Settings{
			Token:  config.BotToken,
			Client: client, // Use custom HTTP client with logging
			URL:    baseURL,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed to create bot: %s", err)
		return
	}

	botSender := b.NewMessageSender()

	bf.CreateFlow(bot, &bs.RefahSettings{
		Sender: botSender,
	})

	log.Println("Bot is running...")
	bot.Start()
}
