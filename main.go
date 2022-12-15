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
	"github.com/joho/godotenv"
	"github.com/maloquacious/moly/engine"
	"log"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("main: %+v\n", err)
	}
	if err := run(); err != nil {
		log.Fatalf("main: %+v\n", err)
	}
}

func run() error {
	e := engine.New()
	return e.Loop()
}
