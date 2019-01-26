// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"bufio"
	"log"
	"math/rand"
	"time"
	"strconv"

	"github.com/spf13/cobra"

	"iquiz"
)

var name string
var reader = bufio.NewReader(os.Stdin)

//http://golangcookbook.blogspot.com/2012/11/generate-random-number-in-given-range.html
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play our amazing quiz game",
	Long: `Play our amazing quiz game`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("* Lets start answering some serious questions 😱\n")
		fmt.Println("* Please Enter your Alter Ego (Name):")

		
		iBuff, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Could not read your Alter Ego: %v\n", err)
		}
		name = strings.TrimSuffix(iBuff, "\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Quiz for %s \n", name)

		// Retrieve all Question from Server
		l, err := client.List(context.Background(), &iquiz.Void{})
		if err != nil {
			log.Fatalf("Could not load questions from server: %v", err)
			os.Exit(1)
		}

		resp := &iquiz.QuizResponse{
			Name: name,
		}

		for _, q := range l.Questions {
			fmt.Printf("%s \n", q.Question )
			q_pos := 1;
			a_pos := random(1, 4)
			is_correct := false
			tmp_answer := ""
			
			for t, a := range q.AnswerOptions {
				fmt.Println( t )
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
		 	ans = strings.TrimSpace(ans)

		 	i, err := strconv.ParseInt(ans, 10, 64)
		 	if err != nil || i < 0 || i > 4 || i != int64(a_pos) {
		 		//fmt.Printf( "* Points lost for incorrect input! *\n")
	 		
			 	// Calculate actual answer:
			 	if i > int64(a_pos) {
			 		log.Print("-- Greater then position of answer!!")
			 		tmp_answer = q.AnswerOptions[i-2]
			 	} else {
			 		tmp_answer = q.AnswerOptions[i-1]
			 	}

		 	} else {
		 		fmt.Printf( "--- CORRECT " )
		 		is_correct = true
		 		tmp_answer = q.CorrectAnswer
		 	}		 	

		 	ans_report := &iquiz.Answer {
		 		Id: q.Id,
		 		Answer: tmp_answer,
		 		IsCorrect: is_correct,
		 	}
		 	resp.Answers = append(resp.Answers, ans_report)
		}

		_, err2 := client.Response(context.Background(), resp)
		if err2 != nil {
			log.Fatalf("Could not send questions to server: %v", err)
			os.Exit(1)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println( "Done already? \n")
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
