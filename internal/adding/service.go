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

package adding

import (
	"time"
)

type Service interface {
	AddAdr(title string, superseded int, link []int) error
}

type Repository interface {
	AddAdr(adr ADR) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) AddAdr(title string, superseded int, link []int) error {
	t := time.Now()
	a := ADR{
		Title:      title,
		Date:       t.Local().Format("2006-01-02"),
		Status:     "Proposed",
		Supersedes: 1,
		Links:      []int{1, 2, 3},
	}
	return s.r.AddAdr(a)
}
