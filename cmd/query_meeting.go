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


// query_meetingCmd represents the query_meeting command
var query_meetingCmd = &cobra.Command{
	Use:   "query_meeting -s [start time] -e [end time]",
	Short: "query meeting",
	Long: `Input command model like : query_meeting -s 2014-11-14/12:00 -e 2014-12-12/12:00

	then we will query the meeting which is taken place between 2014-11-14/12:00 and 2014-12-12/12:00`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_s, _ := cmd.Flags().GetString("start")
		tmp_e, _ := cmd.Flags().GetString("end")
		if service.GetFlag() == true {
			service.Query_meeting(tmp_s, tmp_e)
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(query_meetingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// query_meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// query_meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	query_meetingCmd.Flags().StringP("start", "s", "", "start time")
	query_meetingCmd.Flags().StringP("end", "e", "", "end time")

}
