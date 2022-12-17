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

package engine

import (
	"fmt"
	"log"
)

// Loop is the game loop.
func (e *Engine) Loop() error {
	for tick := 1; tick <= e.ticksPerTurn; tick++ {
		log.Printf("loop: processing tick %d\n", tick)
		if err := e.tick(tick); err != nil {
			return fmt.Errorf("loop: %w", err)
		}
	}
	return nil
}

// tick is one iteration of the game loop
func (e *Engine) tick(tick int) error {
	return nil
}
