package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tidwall/gjson"
	tb "gopkg.in/tucnak/telebot.v2"
)

func getURL(url, scheme string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url = gjson.Get(string(body), scheme).Str
	// log.Println(string(body))
	log.Println(url)
	return url
}

func main() {
	Token := os.Getenv("OhMyRBQBot_TOKEN")
	println(Token)

	b, err := tb.NewBot(tb.Settings{
		Token: Token,
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		// URL: "http://195.129.111.17:8012",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// prpr

	moreprprBtn := tb.InlineButton{
		Unique: "moreprpr",
		Text:   "再整一张😘",
	}

	moreprprKeys := [][]tb.InlineButton{
		[]tb.InlineButton{moreprprBtn},
	}

	b.Handle(&moreprprBtn, func(m *tb.Callback) {
		b.Notify(m.Message.Chat, "typing")
		// b.Notify(update.Message.Chat.ID, tgbotapi.ChatTyping)
		url := getURL("https://api.ixiaowai.cn/api/api.php?return=json", "imgurl")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Message.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moreprprKeys,
		})
	})

	b.Handle("/prpr", func(m *tb.Message) {
		b.Notify(m.Chat, "typing")
		url := getURL("https://api.ixiaowai.cn/api/api.php?return=json", "imgurl")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moreprprKeys,
		})
	})

	// meow

	moremeowBtn := tb.InlineButton{
		Unique: "moremeow",
		Text:   "再整一张😘",
	}

	moremeowKeys := [][]tb.InlineButton{
		[]tb.InlineButton{moremeowBtn},
	}

	b.Handle(&moremeowBtn, func(m *tb.Callback) {
		b.Notify(m.Message.Chat, "typing")
		url := getURL("http://thecatapi.com/api/images/get?format=json&type=jpg,png", "0.url")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Message.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moremeowKeys,
		})
	})

	b.Handle("/meow", func(m *tb.Message) {
		b.Notify(m.Chat, "typing")
		url := getURL("http://thecatapi.com/api/images/get?format=json&type=jpg,png", "0.url")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moremeowKeys,
		})
	})

	// meowmeow

	moremeowmeowBtn := tb.InlineButton{
		Unique: "moremeowmeow",
		Text:   "再整一张😘",
	}

	moremeowmeowKeys := [][]tb.InlineButton{
		[]tb.InlineButton{moremeowmeowBtn},
	}

	b.Handle(&moremeowmeowBtn, func(m *tb.Callback) {
		b.Notify(m.Message.Chat, "typing")
		url := getURL("http://thecatapi.com/api/images/get?format=json&type=gif", "0.url")
		p := &tb.Document{File: tb.FromURL(url)}
		b.Send(m.Message.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moremeowmeowKeys,
		})
	})

	b.Handle("/meowmeow", func(m *tb.Message) {
		b.Notify(m.Chat, "typing")
		url := getURL("http://thecatapi.com/api/images/get?format=json&type=gif", "0.url")
		p := &tb.Document{File: tb.FromURL(url)}
		b.Send(m.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: moremeowmeowKeys,
		})
	})

	// nsfw

	morensfwBtn := tb.InlineButton{
		Unique: "morensfw",
		Text:   "再整一张😘",
	}

	morensfwKeys := [][]tb.InlineButton{
		[]tb.InlineButton{morensfwBtn},
	}

	b.Handle(&morensfwBtn, func(m *tb.Callback) {
		b.Notify(m.Message.Chat, "typing")
		url := getURL("https://yande.re/post.json?tags=order:random&limit=1", "0.file_url")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Message.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: morensfwKeys,
		})
	})

	b.Handle("/nsfw", func(m *tb.Message) {
		b.Notify(m.Chat, "typing")
		url := getURL("https://yande.re/post.json?tags=order:random&limit=1", "0.file_url")
		p := &tb.Photo{File: tb.FromURL(url)}
		b.Send(m.Chat, p, &tb.ReplyMarkup{
			InlineKeyboard: morensfwKeys,
		})
	})

	// visithome

	visithomeBtn := tb.InlineButton{
		Unique: "visithome",
		URL:    "https://idealclover.top",
		Text:   "🏠idealclover.top",
	}

	visithomeKeys := [][]tb.InlineButton{
		[]tb.InlineButton{visithomeBtn},
	}

	b.Handle("/visithome", func(m *tb.Message) {
		b.Notify(m.Chat, "typing")
		// url := getURL("http://thecatapi.com/api/images/get?format=json&type=gif", "0.url")
		p := &tb.Sticker{File: tb.File{FileID: "CAADAgADYFwAAuCjggegjPGxr_HLwRYE"}}
		b.Send(m.Chat, p)
		b.Send(m.Chat, "欢迎造访翠翠的家 ww", &tb.ReplyMarkup{
			InlineKeyboard: visithomeKeys,
		})
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "调教我的姿势不对呦ww你不是我的主人吧")
		p := &tb.Sticker{File: tb.File{FileID: "CAADAgADYFwAAuCjggegjPGxr_HLwRYE"}}
		b.Send(m.Chat, p)
	})

	b.Start()
}
