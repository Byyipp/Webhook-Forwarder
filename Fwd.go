package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var previd = "<Insert the Channel ID from Server>" //check if it in Original channel, this is home wehook for original webhooks


func checkid(id string) bool {
	if strings.Compare(id, previd) == 0 {
		return true
	}
	return false
}

//replaces sensitive information that was covered with spoiler text
// customize the text below to whatever you want!
func check(web []*discordgo.MessageEmbed) {
	for i := range web[0].Fields {
		if strings.Contains(web[0].Fields[i].Value, "||") {
			web[0].Fields[i].Value = "||Caught In 4K||"}} //can customize
	web[0].URL = "https://cdn.discordapp.com/attachments/617974460946972692/858492433141596200/EwOidsXXEAAfWRy.png" //can customize
}

func handMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	//check to make sure it is from the certain channel and that it is a webhook
	if checkid(m.ChannelID) && len(m.Embeds) >= 1 {
		id := m.Embeds

		var routine sync.WaitGroup

		//can add multiple methods here to check for certain webhooks from multiple sources
		routine.Add(1)
		go func() {
			check(id)
			routine.Done()
		}()
		routine.Wait()

		webhook := discordgo.WebhookParams{
			Username:  "Webhook Forwarder", //can customize
			Embeds:    id,
			AvatarURL: "https://cdn.discordapp.com/emojis/660223255646437428.png?v=1", //can customize
		}

		_, _ = s.WebhookExecute("<Insert webhookID from webhook URL>", "<Insert webhook token portion from webhook URL>", true, &webhook)
		// example if webhook URL is https://discord.com/api/webhooks/4984116598654541/3Ifmwkle9vsSvASF2f234klnflk2en0vS_v
		// _, _ = s.WebhookExecute("4984116598654541", "4984116598654541/3Ifmwkle9vsSvASF2f234klnflk2en0vS_v", true, &webhook)
	} //webhook id and token from discord webhook^
		//copy + paste this if you want to forward to multiple discord webhook URLS^
}

const token string = "<Insert Bot Token>" //bot token from developer discord site

var BotID string

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	dg.AddHandler(handMessage)

	_ = dg.Open()

	fmt.Println("Forwarder running..")

	<-make(chan struct{})
	return
}
