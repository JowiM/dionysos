// In memory question list
// Questions from https://hobbylark.com/party-games/The-Funniest-Most-Hilarious-Trivia-Game-Ever-Written-Questions-and-Answers

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

	question3 := &iquiz.Question{
		Id:            3,
		Question:      "Who were the first television couple to be shown in bed together on prime time television?",
		CorrectAnswer: "Fred and Wilma Flinstone",
		AnswerOptions: []string{"Kim Kardashian and Kanye West", "Ozzy and Sharon Osbourne", "Jessica Simpson and Nick Lachey"},
	}

	question4 := &iquiz.Question{
		Id:            4,
		Question:      "More trivia team name triva. Here’s the team name: Stuffing Torrey’s Holes Like Tiger. In what state is Torrey Pines located?",
		CorrectAnswer: "California",
		AnswerOptions: []string{"Washington", "Florida", "Hawaii"},
	}

	question5 := &iquiz.Question{
		Id:            5,
		Question:      "What is the only country through which both the equator and the Tropic of Capricorn pass?",
		CorrectAnswer: "Brazil",
		AnswerOptions: []string{"China", "Russia", "Malta"},
	}

	question6 := &iquiz.Question{
		Id:            6,
		Question:      "In what country might you find the Great Fence?",
		CorrectAnswer: "Australia",
		AnswerOptions: []string{"USA", "Mexico", "South Korea"},
	}

	question7 := &iquiz.Question{
		Id:            7,
		Question:      "Where was the greatest difference between annual high and low temperatures recorded?",
		CorrectAnswer: "Russia",
		AnswerOptions: []string{"Antartica", "Norway", "Sweden"},
	}

	question8 := &iquiz.Question{
		Id:            8,
		Question:      "In which country were bananas first grown?",
		CorrectAnswer: "India",
		AnswerOptions: []string{"Thailand", "Peru", "Vietnam"},
	}

	question9 := &iquiz.Question{
		Id:            9,
		Question:      "Where might one find the Ponte Vecchio?",
		CorrectAnswer: "Florence",
		AnswerOptions: []string{"Venice", "Milan", "Bari"},
	}

	question10 := &iquiz.Question{
		Id:            10,
		Question:      "Which American state has the most earthquakes?",
		CorrectAnswer: "Alaska",
		AnswerOptions: []string{"Florida", "Texas", "Washington"},
	}

	question_list := &iquiz.QuestionList{}
	question_list.Questions = append(question_list.Questions, question)
	question_list.Questions = append(question_list.Questions, question2)
	question_list.Questions = append(question_list.Questions, question3)
	question_list.Questions = append(question_list.Questions, question4)
	question_list.Questions = append(question_list.Questions, question5)
	question_list.Questions = append(question_list.Questions, question6)
	question_list.Questions = append(question_list.Questions, question7)
	question_list.Questions = append(question_list.Questions, question8)
	question_list.Questions = append(question_list.Questions, question9)
	question_list.Questions = append(question_list.Questions, question10)

	return question_list
}
