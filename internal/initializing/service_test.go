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

package initializing_test

import (
	"github.com/fuhrmeistery/adr/internal/initializing"
	"github.com/fuhrmeistery/adr/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const template string = `# {{ .Id }}. {{ .Title }}

Date: {{ .Date }}

## Status

{{ .Status }}
{{ if .Supersedes }}
  * Supersedes {{ .Supersedes -}}
{{ end }}
{{- range .Links }}
  * {{ . }}
{{- end }}

## Context

Record the architectural decisions made on this project.

## Decision

We will use Architecture Decision Records, as described by Michael Nygard in this article: http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions

## Consequences

See Michael Nygard's article, linked above.
`

var _ = Describe("Initialize ADR config", func() {
	It("Should create file with default template and ADR directory", func() {
		r := mock.NewRepository()
		s := initializing.NewService(r)
		err := s.AddConfig("adr")
		if err != nil {
			Fail("Error in Service")
		}
		result := r.GetConfig()
		Expect(result.Directory).To(Equal("adr"))
		Expect(result.Template).To(Equal(template))
	})
})
