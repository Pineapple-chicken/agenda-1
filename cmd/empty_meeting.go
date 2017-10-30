// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"agenda/service"
	"github.com/spf13/cobra"
)

// empty_meetingCmd represents the empty_meeting command
var empty_meetingCmd = &cobra.Command{
	Use:   "empty_meeting",
	Short: "empty all meeting you call",
	Long: `Input command model like : empty_meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if service.GetFlag() == true {
			service.Empty_meeting()
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(empty_meetingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// empty_meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// empty_meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
