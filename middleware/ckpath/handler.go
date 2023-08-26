// moly - a game engine inspired by better games
// Copyright (c) 2023 Michael D Henderson.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package ckpath implements middleware to check a request path for invalid runes.
// Derived from https://github.com/hashicorp/go-cleanhttp/blob/master/handlers.go
package ckpath

import (
	"net/http"
	"unicode"
)

// OnlyPrintableRunes verifies that the request path contains only printable runes.
// If it finds any non-printable runes, it returns a Bad Request response.
func OnlyPrintableRunes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r == nil {
			return
		}
		// return an error if the URL contains any non-printable runes.
		for _, ch := range r.URL.Path {
			if !unicode.IsPrint(ch) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
