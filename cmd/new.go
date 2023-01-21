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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fuhrmeistery/adr/pkg/adr"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new ADR",
	Long:  `Create a new ADR.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		filename := strings.Join(args, "-")

		file, err := os.Create("0001-" + filename + ".md")

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		someTitle := "some-name"
		t := time.Now()

		sup := fmt.Sprintf("[ADR-%d](./%d-%s.md)", supersede, supersede, someTitle)

		a := &adr.ADR{
			Id:         1,
			Title:      title,
			Date:       t.Local().Format("2006-01-02"),
			Status:     "Proposed",
			Supersedes: sup,
			Links: []string{
				"[ADR-0001](./0001-template.md)",
				"[ADR-0002](./0002-template.md)",
				"[ADR-0003](./0003-template.md)",
			},
		}
		adr.CreateAdr(file, a)
	},
}

var links []int
var supersede int

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
	newCmd.Flags().IntVarP(&supersede, "supersede", "s", 0, "Mark an old ADR as superseeded")
}
