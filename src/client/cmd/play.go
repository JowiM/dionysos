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
	"log"
	"bufio"

	"github.com/spf13/cobra"

	"iquiz"
	"client/quiz"
)

var name string


// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play our amazing quiz game",
	Long: `Play our amazing quiz game`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("* Lets start answering some serious questions 😱\n")
		fmt.Println("* Please Enter your Alter Ego (Name):")

		var reader = bufio.NewReader(os.Stdin)
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

		resp := quiz.Play( l, name )

		ranking, err2 := client.Response(context.Background(), resp)
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
}
