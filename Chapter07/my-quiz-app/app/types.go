package app

import "time"

type Answer struct {
	Timestamp          time.Time `json:"timestamp" bson:"timestamp"`
	QuestionID         int       `json:"questionID" bson:"questionID"`
	AnswerIndex        int       `json:"answerIndex" bson:"answerIndex"`
	CorrectAnswerIndex int       `json:"correctAnswerIndex" bson:"correctAnswerIndex"`
	Correct            bool      `json:"correct" bson:"correct"`
}

type Question struct {
	ID                 int      `json:"ID" bson:"ID"`
	Text               string   `json:"Text" bson:"Text"`
	CorrectAnswerIndex int      `json:"CorrectAnswerIndex" bson:"CorrectAnswerIndex"`
	Answers            []string `json:"Answers" bson:"Answers"`
}
