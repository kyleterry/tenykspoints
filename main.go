package main

import (
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/kyleterry/tenykspoints/goleg"
	"github.com/op/go-logging"

	"encoding/json"
	"io/ioutil"
	_ "regexp"
	"time"
)

const (
	// GiftRate really should turn into a formula for inflation. Static points
	// are boring.
	GiftRate = 10

	DataIn  = make(chan []byte, 1000)
	DataOut = make(chan []byte, 1000)
)

var (
	ReAddPoints    []string
	ReRemovePoints []string
	ReListPoints   []string
)

type Config struct {
	TenyksChannel   string
	ServicesChannel string
}

type Points struct {
	Person string
	Points int
}

type meta struct {
	LastGift time.Time
}

var log = logging.MustGetLogger("tenykspoints")

func GiftTimer() {
	//TODO: if service crashes, check time since last gifting so people don't
	// wait more than a week to get their points awarded.
	for {
		select {
		case <-time.After(time.Second * 604800):
			IssueGifts()
		}
	}
}

func IssueGifts() {
	//Function to give everyone GiftRate or so points
}

func AddRegexFilters() {
	// TODO: use config list
}

func IncomingHandler(data []byte) {
	// IncomingHandler will regex data to see if it fits the mold
	// returns fucking nothing
}

func main() {
	var err error
	var input []byte

	input, err = ioutil.ReadFile("/home/kyle/config.json")

	if err != nil {
		log.Fatal(err)
	}

	conf := new(config)
	err = json.Unmarshal(input, &conf)

	go GiftTimer()

	for {
		select {
		case data := <-DataIn:
			go IncomingHandler(data)
		}
	}
}
