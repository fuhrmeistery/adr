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

package mock

import (
	"github.com/fuhrmeistery/adr/internal/adding"
)

type Repository struct {
	adrs []adding.ADR
}

func NewRepository() *Repository {
	return &Repository{[]adding.ADR{}}
}

func (r *Repository) GetADR() []adding.ADR {
	return r.adrs
}

func (r *Repository) AddAdr(a adding.ADR) error {
	r.adrs = append(r.adrs, a)
	return nil
}
