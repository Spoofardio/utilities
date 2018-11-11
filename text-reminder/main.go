package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/sfreiberg/gotwilio"
)

var twilio *gotwilio.Twilio

func main() {
	twilio = gotwilio.NewTwilioClient(sid, token)
	log.Println("Reminder Service Started")
	log.Println(getCatFact())
	gocron.Every(1).Day().At(reminderTime).Do(sendText)
	<-gocron.Start()
}

func sendText() {
	fact := getCatFact()
	twilio.SendSMS(from, to, fact, "", "")
	log.Println("Message Sent : ", fact)
}

type catFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func getCatFact() string {
	response, err := http.Get("https://catfact.ninja/fact?max_length=115")
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return message
	} else {
		var fact catFact
		data, _ := ioutil.ReadAll(response.Body)

		err := json.Unmarshal([]byte(data), &fact)
		if err != nil {
			log.Printf("Json Unmarshal error: %s\n", err)
			return message
		}
		return fact.Fact + " " + message
	}
}
