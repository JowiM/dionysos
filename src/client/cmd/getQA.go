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
	"fmt"
	"os"
	"context"
	"log"
	
	"github.com/spf13/cobra"
	"github.com/kyokomi/emoji"

	"iquiz"
)

// getQACmd represents the getQA command
var getQACmd = &cobra.Command{
	Use:   "getQA",
	Short: "List all questions and its corresponding answers",
	Long: `List all questions and corresponding answers aka Cheat Sheet.`,
	Run: func(cmd *cobra.Command, args []string) {
		emoji.Println("--- :disappointed: CheatSheet :disappointed: ---")

		l, err := client.List(context.Background(), &iquiz.Void{})
		if err != nil {
			log.Fatalf("Could not get questions and Answers: %v", err)
			os.Exit(1)
		}

		for _, t := range l.Questions {
			fmt.Printf("ID: %d \n - Question: %s \n - Answer: %s \n", t.Id, t.Question, t.CorrectAnswer)
		}
	},
}

func init() {
	rootCmd.AddCommand(getQACmd)
}
