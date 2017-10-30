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
	"agenda/service"
	"github.com/spf13/cobra"
)



// log_inCmd represents the log_in command
var log_inCmd = &cobra.Command{
	Use:   "log_in -n [username] -p [password]",
	Short: "log in",
	Long: `Input command mode like : log_in -n Go -p 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		service.Log_in(tmp_n, tmp_p)
	},
}

func init() {
	RootCmd.AddCommand(log_inCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// log_inCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// log_inCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	log_inCmd.Flags().StringP("name", "n", "", "user name")
	log_inCmd.Flags().StringP("password", "p", "", "user password")
	
	/*if () {
		
		
		
		
		
		
		


	}*/
}
