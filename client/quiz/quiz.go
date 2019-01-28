package quiz

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strings"
	"strconv"
	"bufio"

	"iquiz"
)

type Quiz struct {
	AllQuestions func()
	Play func( q_list *iquiz.QuestionList, name string ) ( *iquiz.QuizResponse )
}

//http://golangcookbook.blogspot.com/2012/11/generate-random-number-in-given-range.html
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func verify_answer( ans string, answer_pos int, ans_options []string, correct_ans string ) ( answered string, is_correct bool ){
	// Convert to integer
	ans = strings.TrimSpace(ans)
	i, err := strconv.ParseInt(ans, 10, 64)
	// @TODO update to case
	// Not Integer
	if err != nil {
		return "", false
	} else if i < 1 || i > 4 { // Not In range
	 	return "", false
	} else if i != int64(answer_pos) { // Not Correct but in range
	 	// Calculate actual answer if before answer:
	 	if i > int64(answer_pos) {
	 		i = i -1
	 	} 
	 	i = i-1

		return ans_options[i], false
	} else { // Correct
		return correct_ans, true
 	}		 	
}

func Play( q_list *iquiz.QuestionList, name string ) ( *iquiz.QuizResponse ) {
	var reader = bufio.NewReader(os.Stdin)

	resp := &iquiz.QuizResponse{
		Name: name,
	}

	for _, q := range q_list.Questions {
			// Answer position
			q_pos := 1;
			// Randomly generate anser position
			a_pos := random(1, 4)

			fmt.Printf("%s \n", q.Question )
			
			for _, a := range q.AnswerOptions {
				if( a_pos == q_pos ) {
					fmt.Printf("%d. %s \n", q_pos, q.CorrectAnswer)
					q_pos++					
				}

				fmt.Printf("%d. %s \n", q_pos, a)
				q_pos++				
			}

			ans, err := reader.ReadString('\n')
		 	if err != nil {
		  		fmt.Fprintln(os.Stderr, err)
		 	}
		 	
		 	tmp_answer, is_correct := verify_answer( ans, a_pos, q.AnswerOptions, q.CorrectAnswer )
		 	ans_report := &iquiz.Answer {
		 		Id: q.Id,
		 		Answer: tmp_answer,
		 		IsCorrect: is_correct,
		 	}
		 	resp.Answers = append(resp.Answers, ans_report)
		}

		return resp
}