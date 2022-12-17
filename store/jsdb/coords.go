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

package jsdb

import (
	"fmt"
)

// CoordsList is a list of coordinates
type CoordsList []*Coords

func (c CoordsList) Len() int {
	return len(c)
}

func (c CoordsList) Less(i, j int) bool {
	if c[i].Q < c[j].Q {
		return true
	} else if c[i].Q == c[j].Q {
		return c[i].R < c[j].R
	}
	return false
}

func (c CoordsList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Coords are the hex-grid coordinates
type Coords struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

func (c Coords) Less(a Coords) bool {
	if c.Q < a.Q {
		return true
	} else if c.Q == a.Q {
		return c.R < a.R
	}
	return false
}

// String implements the Stringer interface
func (c Coords) String() string {
	return fmt.Sprintf("%d,%d", c.Q, c.R)
}
