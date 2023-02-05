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
	"errors"
	"html/template"
	"log"
	"os"
	"sync"

	"github.com/fuhrmeistery/adr/internal/adding"
)

//go:embed template.md
var content embed.FS

type Storage struct {
	m         sync.Mutex
	directory string
}

func NewStorage(directory string) *Storage {
	return &Storage{directory: directory}
}

func (s *Storage) AddAdr(a adding.ADR) error {
	s.m.Lock()
	defer s.m.Unlock()
	adr := s.createADR(a)
	return s.saveADR(adr)
}

func (s *Storage) createADR(a adding.ADR) ADR {
	id, err := s.getNextADRId()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot get next Id")
	}
	superseded, err := s.createLinkToSuperseded(a.Supersedes)

	links := []string{
		CreateLink("0002-something-else.md"),
	}
	return ADR{
		Id:         id,
		Title:      a.Title,
		Date:       a.Date,
		Status:     a.Status,
		Supersedes: superseded,
		Links:      links,
	}
}

func (s *Storage) createLinkToSuperseded(superseded int) (string, error) {
	filename, err := s.getFilenameById(superseded)
	if err != nil {
		return "", err
	}
	return CreateLink(filename), nil
}

func (s *Storage) getFilenameById(id int) (string, error) {
	files, err := os.ReadDir(s.directory)
	if err != nil {
		return "", err
	}
	if len(files) < id || id == 0 {
		return "", errors.New("id does not exist")
	}
	return files[id-1].Name(), nil
}

func (s *Storage) getNextADRId() (int, error) {
	files, err := os.ReadDir(s.directory)
	if err != nil {
		return 0, err
	}
	return len(files) + 1, nil
}

func (s *Storage) saveADR(a ADR) error {
	filename := CreateFilename(a.Id, a.Title)
	wr, err := os.Create(s.directory + "/" + filename)

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
	return tmpl.Execute(wr, a)

}
