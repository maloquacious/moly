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
	"bytes"
	"fmt"
)

type svg struct {
	id      string
	viewBox struct {
		minX, minY    float64
		width, height float64
	}
	circles  []*circle
	fontSize int
}

type circle struct {
	col, row       int     // coordinates of the hex
	cx, cy, radius float64 // center of the hex
	style          struct {
		fill        string
		stroke      string
		strokeWidth string
	}
	text string
}

func (tiles Tiles) ToSvg() *svg {
	size := 12.0
	width, height := 2*size, 2*size
	radius := size * 0.88

	s := svg{id: "worldmap", fontSize: 14}

	// create the circles
	for _, t := range tiles {
		offset := 0.0
		if t.Col%2 == 1 {
			offset += height / 2
		}
		c := &circle{
			row:    t.Row,
			col:    t.Col,
			cy:     height + float64(t.Row)*height + offset,
			cx:     width + float64(t.Col)*width,
			radius: radius,
			text:   string(t.SaveChar),
		}
		c.style.fill = t.Color.ToFill()
		c.style.stroke = "grey"
		c.style.strokeWidth = "2px"

		s.circles = append(s.circles, c)
		if s.viewBox.height < c.cy {
			s.viewBox.height = c.cy
		}
		if s.viewBox.width < c.cx {
			s.viewBox.width = c.cx
		}
	}
	s.viewBox.height += height
	s.viewBox.width += width

	return &s
}

func (s *svg) Bytes() []byte {
	b := &bytes.Buffer{}

	b.WriteString("<svg")
	if s.id != "" {
		b.WriteString(fmt.Sprintf(" id=%q", s.id))
	}
	b.WriteString(fmt.Sprintf(` width="%f" height="%f"`, s.viewBox.width+40, s.viewBox.height+40))
	b.WriteString(fmt.Sprintf(` viewBox="%f %f %f %f"`, s.viewBox.minX, s.viewBox.minY, s.viewBox.width+40, s.viewBox.height+40))
	b.WriteString(` xmlns="http://www.w3.org/2000/svg">`)
	b.WriteByte('\n')

	for _, c := range s.circles {
		b.WriteString(fmt.Sprintf(`<circle cx="%f" cy="%f" r="%f" fill=%q stroke=%q stroke-width=%q />`, c.cx, c.cy, c.radius, c.style.fill, c.style.stroke, c.style.strokeWidth))
		b.WriteByte('\n')

		//if c.text != "" {
		//	yOffset := float64(s.fontSize) * -0.2
		//	b.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="black" font-size="%d" font-weight="bold">%s</text>`, c.cx, c.cy-yOffset, s.fontSize, c.text))
		//	b.WriteByte('\n')
		//}
	}

	b.WriteString("</svg>\n")

	return b.Bytes()
}
