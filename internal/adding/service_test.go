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

package adding_test

import (
	"time"

	"github.com/fuhrmeistery/adr/internal/adding"
	"github.com/fuhrmeistery/adr/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ADR Adding", func(title string, superseded int, links []int, expected adding.ADR) {
	r := mock.NewRepository()
	s := adding.NewService(r)
	s.AddAdr(title, superseded, links)
	result := r.Get()[0]
	Expect(result).To(Equal(expected))

},
	Entry("Should Add ADR", "Long Title",
		1,
		[]int{1, 2, 3},
		adding.ADR{
			Title:      "Long Title",
			Date:       time.Now().Local().Format(("2006-01-02")),
			Status:     "Proposed",
			Supersedes: 1,
			Links:      []int{1, 2, 3},
		},
	),
)
