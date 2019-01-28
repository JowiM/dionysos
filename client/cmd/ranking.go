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
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/JowiM/dionysus/iquiz"
)

// rankingCmd represents the ranking command
var rankingCmd = &cobra.Command{
	Use:   "ranking",
	Short: "List the current rankings",
	Long: `List all the rankings`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("** Ranking called **")

		// Fetch all Rankings 
		rankings, err := client.Rankings(context.Background(), &iquiz.Void{})
		if err != nil {
			log.Fatalf("Could not load rankings from server: %v", err)
			os.Exit(1)
		}

		// List all Rankings
		for i, rank := range rankings.Rankings {
			fmt.Printf( "%d - Point: %d - { Name: %s } \n", i, rank.Points, rank.Name )
		}
	},
}

func init() {
	rootCmd.AddCommand(rankingCmd)
}
