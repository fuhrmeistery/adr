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
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CreateLink(filename string) string {
	id, err := strconv.Atoi(strings.Split(filename, "-")[0])
	if err != nil {
		log.Fatal("Provided Filename is invalid")
	}
	return fmt.Sprintf("[ADR-%d](./%s)", id, filename)
}

func CreateFilename(id int, title string) string {
	sId := strconv.Itoa(id)
	for len(sId) < 4 {
		sId = "0" + sId
	}
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	return fmt.Sprintf("%s-%s.md", sId, title)
}
