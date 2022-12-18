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
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/maloquacious/moly/enums"
	"golang.org/x/image/font/gofont/goregular"
)

func (tiles Tiles) ToPNG(name string) error {
	dc := gg.NewContext(6800, 7800)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	face := truetype.NewFace(font, &truetype.Options{Size: 14})
	dc.SetFontFace(face)

	// create the circles
	for _, t := range tiles {
		// cy and cx are the center point of the cell
		cx, cy := t.centerPoint()
		dc.SetColor(t.Color.ToRGB())
		dc.DrawRegularPolygon(6, cx, cy, TILESIZE*0.88, 0)
		dc.Fill()
		if t.Terrain != enums.TerrOcean {
			dc.SetRGB(0, 0, 0)
			dc.DrawStringAnchored(fmt.Sprintf("%d %d", t.Col, t.Row), cx, cy-0.375*TILEHEIGHT, 0.5, 0.5)
		}
	}
	return dc.SavePNG(name)
}
