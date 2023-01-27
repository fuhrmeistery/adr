/*
Copyright © 2023 Yannik Fuhrmeister <yannik.fuhrmeister@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"
	"strings"

	"github.com/fuhrmeistery/adr/internal/adding"
	"github.com/fuhrmeistery/adr/internal/storage/filesystem"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new ADR",
	Long:  `Create a new ADR.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repository := filesystem.NewStorage()
		service := adding.NewService(repository)
		title := strings.Join(args, " ")
		err := service.AddAdr(title, superseded, links)
		if err != nil {
			log.Fatalln(err)
		}

	},
}

var links []int
var superseded int

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().IntSliceVarP(&links, "link", "l", nil, "Create a link between the new ADR and an existing one")
	newCmd.Flags().IntVarP(&superseded, "supersede", "s", 0, "Mark an old ADR as superseeded")
}
