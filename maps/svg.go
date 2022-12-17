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
	"math"
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
	points []point
	text   string
}

type point struct {
	x, y float64
}

func (tiles Tiles) ToSvg(hexes, mono bool) *svg {
	width, height := 2*TILESIZE, math.Sqrt(3)*TILESIZE

	radius := TILESIZE * 0.88

	s := svg{id: "worldmap", fontSize: 14}

	var points []point
	for theta := 0.0; theta < math.Pi*2.0; theta += math.Pi / 3.0 {
		points = append(points, point{x: radius * math.Cos(theta), y: radius * math.Sin(theta)})
	}
	// create the circles
	for _, t := range tiles {
		// cy and cx are the center point of the cell
		cx, cy := t.centerPoint()

		c := &circle{
			row:    t.Row,
			col:    t.Col,
			cy:     cy,
			cx:     cx,
			radius: radius,
			text:   string(t.SaveChar),
		}
		if mono {
			c.style.fill = "none"
		} else {
			c.style.fill = t.Color.ToFill()
		}
		c.style.stroke = "grey"
		c.style.strokeWidth = "2px"

		if hexes {
			for _, p := range points {
				px, py := c.cx+p.x, c.cy+p.y
				c.points = append(c.points, point{x: px, y: py})
				if s.viewBox.height < py {
					s.viewBox.height = py
				}
				if s.viewBox.width < px {
					s.viewBox.width = px
				}
			}
		}

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

func (s *svg) Bytes(noScale bool) []byte {
	b := &bytes.Buffer{}

	b.WriteString("<svg")
	if s.id != "" {
		b.WriteString(fmt.Sprintf(" id=%q", s.id))
	}
	if !noScale {
		b.WriteString(fmt.Sprintf(` width="%f" height="%f"`, s.viewBox.width+TILESIZE, s.viewBox.height+TILESIZE))
	}
	b.WriteString(fmt.Sprintf(` viewBox="%f %f %f %f"`, s.viewBox.minX, s.viewBox.minY, s.viewBox.width, s.viewBox.height))
	b.WriteString(` xmlns="http://www.w3.org/2000/svg">`)
	b.WriteByte('\n')

	for _, c := range s.circles {
		if len(c.points) == 0 {
			b.WriteString(fmt.Sprintf(`<circle cx="%f" cy="%f" r="%f" fill=%q stroke=%q stroke-width=%q />`, c.cx, c.cy, c.radius, c.style.fill, c.style.stroke, c.style.strokeWidth))
		} else {
			b.WriteString(fmt.Sprintf(`<polygon fill=%q  stroke=%q stroke-width=%q points="`, c.style.fill, c.style.stroke, c.style.strokeWidth))
			for _, p := range c.points {
				b.WriteString(fmt.Sprintf("%f,%f ", p.x, p.y))
			}
			b.WriteString(`"></polygon>`)
		}
		b.WriteByte('\n')

		if c.text != "" {
			yOffset := float64(s.fontSize) * -0.2
			b.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="black" font-size="%d" font-weight="bold">%s</text>`, c.cx, c.cy-yOffset, s.fontSize, c.text))
			b.WriteByte('\n')
		}
	}

	b.WriteString("</svg>\n")

	return b.Bytes()
}
