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

package filesystem_test

import (
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

var _ = DescribeTable("Markdown Filename", func(id int, title string, expected string) {
},
	Entry("Should pad with zeros", 1, "Some Very Long Title", "0001-some-very-long-title.md"),
	Entry("Should not pad with zeros", 1000, "Short Title", "1000-short-title.md"),
)

var _ = Describe("Markdown CreateLink", func() {
	It("Should create markdown link", func() {
	})
})
