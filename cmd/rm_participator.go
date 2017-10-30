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


// rm_participatorCmd represents the rm_participator command
var rm_participatorCmd = &cobra.Command{
	Use:   "rm_participator -n [username] -t [meeting title]",
	Short: "remove a participator from a meeting",
	Long: `Input command mode like : rm_participator -n Go -t BeautifulGo

	then if the participator and meeting exist, then we will remove Go from meeting BeautifulGo`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_t, _ := cmd.Flags().GetString("title")
		if service.GetFlag() == true {
			service.Rm_participator(tmp_n, tmp_t)
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(rm_participatorCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rm_participatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rm_participatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rm_participatorCmd.Flags().StringP("name", "n", "", "user name")
	rm_participatorCmd.Flags().StringP("title",  "t", "", "meeting title")

}
