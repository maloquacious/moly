/*
 * moly - a game engine inspired by better games
 * Copyright (C) 2022 Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"github.com/playbymail/moly/cli"
	"github.com/playbymail/moly/engine"
	"github.com/playbymail/moly/pkg/dot"
	"github.com/playbymail/moly/store/jsdb"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	started := time.Now()

	rand.Seed(started.UnixNano())

	log.SetFlags(log.LstdFlags | log.LUTC)

	if err := dot.Load("MOLY", true, true); err != nil {
		log.Fatalf("main: %+v\n", err)
	}

	rv := 0
	if err := cli.Execute(); err != nil {
		log.Printf("\n%+v\n", err)
		rv = 2
	}

	log.Printf("\n")
	log.Printf("completed in %v\n", time.Now().Sub(started))

	os.Exit(rv)
}

func run() error {
	data, err := os.ReadFile("worldmap.json")
	if err != nil {
		return err
	}
	m, err := jsdb.MapLoad(data)
	if err != nil {
		return err
	}
	for k, v := range m {
		log.Printf("%s: %v\n", k, v)
	}

	e := engine.New()
	return e.Loop()
}
