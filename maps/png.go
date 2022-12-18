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

import (
	"github.com/fogleman/gg"
)

func (tiles Tiles) ToPNG(name string) error {
	dc := gg.NewContext(6800, 7800)
	//dc.SetRGB(0, 0, 0)
	dc.SetRGB(0x1e/ 255.0, 0x90/ 255.0, 0xff/ 255.0)
	dc.Clear()

	// create the circles
	for _, t := range tiles {
		r,g,b := t.Color.ToRGB()
		dc.SetRGB(r / 255, g  / 255, b  / 255)
		// cy and cx are the center point of the cell
		cx, cy := t.centerPoint()
		dc.DrawRegularPolygon(6, cx, cy, TILESIZE*0.88, 0)
		dc.DrawRegularPolygon(6, cx, cy, TILESIZE, 0)
		dc.Fill()
	}
	return dc.SavePNG(name)
}
