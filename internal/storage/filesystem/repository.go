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

package filesystem

import (
	"embed"
	"html/template"
	"os"

	"github.com/fuhrmeistery/adr/internal/adding"
)

//go:embed template.md
var content embed.FS

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) AddAdr(a adding.ADR) error {
	superseded := "0001-some-name.md"
	links := []string{
		CreateLink("0002-something-else.md"),
	}
	adr := ADR{
		Id:         1,
		Title:      a.Title,
		Date:       a.Date,
		Status:     a.Status,
		Supersedes: CreateLink(superseded),
		Links:      links,
	}

	wr, err := os.Create("0001-" + a.Title + ".md")

	if err != nil {
		return err
	}
	defer wr.Close()

	tmp, err := content.ReadFile("template.md")
	if err != nil {
		return err
	}

	tmpl, err := template.New("ADR").Parse(string(tmp))
	if err != nil {
		return err
	}
	return tmpl.Execute(wr, adr)
}
