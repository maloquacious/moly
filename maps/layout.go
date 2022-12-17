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

package maps

import "math"

const (
	TILEOFFSET = -1 // odd rows
	TILESIZE   = 44.0
	TILESIZEX  = TILESIZE
	TILESIZEY  = TILESIZE
	TILEWIDTH  = 2 * TILESIZE
)

var (
	TILEHEIGHT = math.Sqrt(3) * TILESIZE
	layout     = Orientation{ // flat layout
		f0: 3.0 / 2.0, f1: 0.0, f2: math.Sqrt(3.0) / 2.0, f3: math.Sqrt(3.0),
		b0: 2.0 / 3.0, b1: 0.0, b2: -1.0 / 3.0, b3: math.Sqrt(3.0) / 3.0,
		start_angle: 0.0,
	}
)

// Orientation stores the forward matrix (the fN variables) and backward matrix
// (the bN variables), plus the start angle. The start angle determines if we
// have a "flat top" (which is 0°) or "pointy top" (which is 60°) hex.
type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	// The starting angle should be 0.0 for 0° (flat top) or 0.5 for 60° (pointy top).
	start_angle float64 // in multiples of 60°
}

// centerPoint returns the center point of the tile on the screen.
func (t *Tile) centerPoint() (x, y float64) {
	M := layout
	q, r := t.Col, t.Row-(t.Col+int(TILEOFFSET)*(t.Col&1))/2
	//s := -q - r

	return (M.f0*float64(q)+M.f1*float64(r))*TILESIZEX + TILEWIDTH, (M.f2*float64(q)+M.f3*float64(r))*TILESIZEY + TILEHEIGHT
}
