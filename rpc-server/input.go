package main

import (
	"fmt"
	"strings"
)

type Input struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// ID will be the user name of sender and receiver.
func GetUniqueID(chat string) string {
	splitChat := strings.Split(strings.ToLower(chat), ":")
	participantOne, participantTwo := splitChat[0], splitChat[1]

	var uniqueID string
	if participantOne > participantTwo {
		uniqueID = fmt.Sprintf("%s:%s", participantOne, participantTwo)
	} else {
		uniqueID = fmt.Sprintf("%s:%s", participantTwo, participantOne)
	}
	return uniqueID
}