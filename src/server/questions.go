package main

import (
	"iquiz"
)

type Questions struct {
}

func (q Questions) GetQuestions() *iquiz.QuestionList {
	question := &iquiz.Question{
		Id:            1,
		Question:      "Who is the only member of ZZ Top who doesn’t have a beard?",
		CorrectAnswer: "Frank Beard",
		AnswerOptions: []string{"Jimmy Page", "Frank Zampa", "Dusty Hill"},
	}

	question2 := &iquiz.Question{
		Id:            2,
		Question:      "Who was the first Twitter user to reach 20 million followers? Her songs include The Edge of Glory, Judas, Born This Way, Bad Romance, and Poker Face?",
		CorrectAnswer: "Lady Gaga",
		AnswerOptions: []string{"Jay-Z", "Beyonce", "Judas Priest"},
	}

	question_list := &iquiz.QuestionList{}
	question_list.Questions = append(question_list.Questions, question)
	question_list.Questions = append(question_list.Questions, question2)

	return question_list
}
