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

// Package static implements a static file server as middleware
package static

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Static middleware serves static files
func Static(root string) func(http.Handler) http.Handler {
	// convert the root to an absolute path in case the caller changes directories on us
	var err error
	root, err = filepath.Abs(root)
	if err != nil {
		panic(err)
	}
	log.Printf("[static] root %s\n", root)

	files := make(map[string]*cache)
	for _, path := range []string{
		"browserconfig.xml",
		"favicon.ico",
		"humans.txt",
		"icon.png",
		"robots.txt",
		"rules.html",
		"site.webmanifest",
		"tile-wide.png",
		"tile.png",
		filepath.Join("css", "main.css"),
		filepath.Join("css", "normalize.css"),
		filepath.Join("js", "htmx.min-1.9.4.js"),
	} {
		file := filepath.Clean(filepath.Join(root, path))
		cf := cacheFile(filepath.Clean(filepath.Join(root, path)))
		if cf == nil {
			log.Printf("[static] missed %q\n", path)
			continue
		}
		files[filepath.Base(file)] = cf
		log.Printf("[static] cached %q\n", path)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c, ok := files[filepath.Base(r.URL.Path)]; ok {
				if match := r.Header.Get("If-None-Match"); match != "" {
					if strings.Contains(match, c.etag) {
						log.Printf("[static] %s %q: e-tag cached\n", r.Method, r.URL.Path)
						w.WriteHeader(http.StatusNotModified)
						return
					}
				}
				log.Printf("[static] %s %q: from cache\n", r.Method, r.URL.Path)
				w.Header().Set("Content-Type", c.contentType)
				w.Header().Set("Etag", c.etag)
				w.Header().Set("Cache-Control", c.maxAge)
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(c.b)
				return
			}

			file := root + path.Clean(r.URL.Path)
			log.Printf("[static] file %q\n", file)

			stat, err := os.Stat(file)
			if err != nil || stat.IsDir() || !stat.Mode().IsRegular() {
				// we want to serve only regular files, so let the next handler serve the route
				next.ServeHTTP(w, r)
				return
			}

			// open it and defer the close so that we don't leak resources.
			rdr, err := os.Open(file)
			if err != nil {
				log.Printf("[static] error %q: %v\n", file, err)
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			defer rdr.Close()

			// and serve it
			http.ServeContent(w, r, file, stat.ModTime(), rdr)
		})
	}
}
