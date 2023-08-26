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

// Package static implements a static file server with cacheing.
package static

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Handler serves static files with maybe a cache, too.
func Handler(root string, files ...string) http.HandlerFunc {
	// convert the root to an absolute path in case the caller changes directories on us
	var err error
	root, err = filepath.Abs(root)
	if err != nil {
		panic(err)
	}
	//log.Printf("[static] root %s\n", root)

	fmap := make(map[string]*cache)
	for _, path := range files {
		file := filepath.Clean(filepath.Join(root, path))
		cf := cacheFile(filepath.Clean(filepath.Join(root, path)))
		if cf == nil {
			log.Printf("[static] missed %q\n", path)
			continue
		}
		fmap[filepath.Base(file)] = cf
		log.Printf("[static] cached %q\n", path)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if c, ok := fmap[filepath.Base(r.URL.Path)]; ok {
			if match := r.Header.Get("If-None-Match"); match != "" {
				if strings.Contains(match, c.etag) {
					//log.Printf("[static] %s %q: e-tag cached\n", r.Method, r.URL.Path)
					w.WriteHeader(http.StatusNotModified)
					return
				}
			}
			//log.Printf("[static] %s %q: from cache\n", r.Method, r.URL.Path)
			w.Header().Set("Content-Type", c.contentType)
			w.Header().Set("Etag", c.etag)
			w.Header().Set("Cache-Control", c.maxAge)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(c.b)
			return
		}

		file := root + path.Clean(r.URL.Path)
		//log.Printf("[static] file %q\n", file)

		// we want to serve only regular files
		stat, err := os.Stat(file)
		if err != nil || stat.IsDir() || !stat.Mode().IsRegular() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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
	}
}
