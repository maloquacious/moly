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

package cli

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/playbymail/moly/middleware/ckpath"
	"github.com/playbymail/moly/server"
	"github.com/playbymail/moly/server/htmx"
	"github.com/playbymail/moly/service/static"
	"github.com/spf13/cobra"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

// cmdServe runs the serve command
var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "run web server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("[serve] host   %q\n", argsServe.host)
		log.Printf("[serve] port   %q\n", argsServe.port)
		log.Printf("[serve] public %q\n", argsServe.public)

		s := &http.Server{
			Addr:    net.JoinHostPort(argsServe.host, argsServe.port),
			Handler: server.Router(),
		}

		// serverCh := make(chan struct{})
		// go func() {
		//	log.Printf("[INFO] server is listening on %s\n", ":8080")
		//	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		//		log.Fatalf("[ERR] server exited with: %s", err)
		//	}
		//	close(serverCh)
		// }()
		// signalCh := make(chan os.Signal, 1)
		// signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
		// // Wait for interrupt
		// <-signalCh
		// log.Printf("[INFO] received interrupt, shutting down...")
		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()
		// if err := server.Shutdown(ctx); err != nil {
		//	log.Fatalf("[ERR] failed to shutdown server: %s", err)
		// }
		// // If we got this far, it was an interrupt, so don't exit cleanly
		// os.Exit(2)

		r := chi.NewRouter()
		r.Use(ckpath.OnlyPrintableRunes)
		if argsServe.logRequests {
			r.Use(middleware.Logger)
		}
		r.Use(middleware.Heartbeat("/ping"))
		r.Use(middleware.Recoverer)

		// static files
		r.Get("/browserconfig.xml", static.Handler(argsServe.public, "browserconfig.xml"))
		r.Get("/favicon.ico", static.Handler(argsServe.public, "favicon.ico"))
		r.Get("/humans.txt", static.Handler(argsServe.public, "humans.txt"))
		r.Get("/icon.png", static.Handler(argsServe.public, "icon.png"))
		r.Get("/robots.txt", static.Handler(argsServe.public, "robots.txt"))
		r.Get("/rules.html", static.Handler(argsServe.public, "rules.html"))
		r.Get("/site.webmanifest", static.Handler(argsServe.public, "site.webmanifest"))
		r.Get("/tile-wide.png", static.Handler(argsServe.public, "tile-wide.png"))
		r.Get("/tile.png", static.Handler(argsServe.public, "tile.png"))
		r.Get("/css/*", static.Handler(filepath.Join(argsServe.public, "css"), "main.css", "normalize.css"))
		r.Get("/js/*", static.Handler(filepath.Join(argsServe.public, "js"), "htmx.min-1.9.4.js"))

		// mount the api router
		r.Mount("/api", htmx.Router())

		r.Get("/", getIndex())

		s.Handler = r

		_ = http.ListenAndServe(s.Addr, s.Handler)
	},
}

var argsServe = struct {
	public      string // path to public (static) files
	host        string
	port        string
	logRequests bool
}{
	port: "3000",
}

func getIndex() http.HandlerFunc {
	data := []byte(`
<script src="/js/htmx.min-1.9.4.js"></script>
<button hx-post="/api/clicked" hx-swap="outerHTML">Click Me</button>
<div hx-put="/api/messages">Put To Messages</div>`)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content", "text/html")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}
