package main

import (
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/sfreiberg/gotwilio"
)

var twilio *gotwilio.Twilio

func main() {
	twilio = gotwilio.NewTwilioClient(sid, token)
	log.Println("Reminder Service Started : ", time.Now().Format("2006-Jan-02"))
	gocron.Every(1).Day().At(reminderTime).Do(sendText)
	<-gocron.Start()
}

func sendText() {

	twilio.SendSMS(from, to, message, "", "")
	log.Println("Message Sent : ", time.Now().Format("2006-Jan-02"))
}
