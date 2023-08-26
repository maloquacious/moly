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

package static

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// cache files that we prize
type cache struct {
	b           []byte
	etag        string
	contentType string
	maxAge      string
}

func cacheFile(file string) *cache {
	var contentType string
	switch filepath.Ext(file) {
	case ".css":
		contentType = "text/css; charset=utf-8"
	case ".gif":
		contentType = "image/gif"
	case ".html":
		contentType = "text/html"
	case ".ico":
		contentType = "image/x-icon"
	case ".js":
		contentType = "text/javascript"
	case ".png":
		contentType = "image/png"
	case ".txt":
		contentType = "text/plain; charset=utf-8"
	case ".webmanifest":
		contentType = "application/json; charset=utf-8"
	case ".xml":
		contentType = "text/xml"
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	// calculate the SHA-1 hash of the file to use as the e-tag
	h := sha1.New()
	h.Write(data)
	etag := "moly:" + base64.URLEncoding.EncodeToString(h.Sum(nil))

	return &cache{
		contentType: contentType,
		b:           data,
		etag:        etag,
		maxAge:      fmt.Sprintf("max-age=%d", 4*7*24*60*60), // 4 weeks
	}
}
