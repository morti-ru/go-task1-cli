/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// questionCmd represents the question command
var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Parameter for ask question",
	Long: `Parameter for ask question big`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Question asked:", args[0])
		switch args[0] {
			case "How are you?":
				fmt.Println("I am fine")
			case "Who am i?":
				fmt.Println("I am batman!")
			case "Good buy!":
				fmt.Println("See you later!")
			default:
				fmt.Println("I have no answere!")
		}
	},
}

func init() {
	rootCmd.AddCommand(questionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
