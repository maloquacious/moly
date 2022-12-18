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

package enums

import "image/color"

type Color int

const (
	ColorNone          Color = iota
	ColorOcean               = 1
	ColorOceanTwo            = 2
	ColorOceanThree          = 3
	ColorOceanFour           = 4
	ColorPlain               = 5
	ColorPlainTwo            = 6
	ColorDesert              = 7
	ColorDesertTwo           = 8
	ColorMountain            = 9
	ColorMountainTwo         = 10
	ColorSwamp               = 11
	ColorSwampTwo            = 12
	ColorForest              = 13
	ColorForestTwo           = 14
	ColorMountainThree       = 16
	ColorForestThree         = 19
	ColorOrange              = -1
)

func (c Color) ToFill() string {
	switch c {
	case ColorOcean:
		return "darkblue"
	case ColorOceanTwo:
		return "cornflowerblue"
	case ColorOceanThree:
		return "lightsteelblue"
	case ColorOceanFour:
		return "blue"
	case ColorPlain:
		return "navajowhite"
	case ColorPlainTwo:
		return "brown"
	case ColorDesert:
		return "sandybrown"
	case ColorDesertTwo:
		return "khaki"
	case ColorMountain:
		return "sienna"
	case ColorMountainTwo:
		return "saddlebrown"
	case ColorSwamp:
		return "yellowgreen"
	case ColorSwampTwo:
		return "seagreen"
	case ColorForest:
		return "forestgreen"
	case ColorForestTwo:
		return "darkolivegreen"
	case ColorMountainThree:
		return "peru"
	case ColorForestThree:
		return "olive"
	case ColorOrange:
		return "orange"
	}
	return "aliceblue"
}

var (
	cAliceBlue = color.RGBA{R: 0xf0, G: 0xf8, B: 0xff} // alice blue
	cDesert    = color.RGBA{R: 0xfa, G: 0xa4, B: 0x60} // sandy brown
	cForest    = color.RGBA{R: 0x22, G: 0x8b, B: 0x22} // forest green
	cMountain  = color.RGBA{R: 0xa0, G: 0x52, B: 0x2d} // sienna
	cOceanBlue = color.RGBA{R: 0x41, G: 0x69, B: 0xe1} // royal blue
	cOlive     = color.RGBA{R: 0x80, G: 0x80, B: 0x00} // olive
	cOrange    = color.RGBA{R: 0xff, G: 0xa5, B: 0x00} // orange
	cPlain     = color.RGBA{R: 0xff, G: 0xde, B: 0xad} // navajo white
	cSixteen   = color.RGBA{R: 0xcd, G: 0x85, B: 0x3f} // peru
	cSwamp     = color.RGBA{R: 0x9a, G: 0xcd, B: 0x32} // yellow green
)

func (c Color) ToRGB() color.Color {
	switch c {
	case ColorOcean:
		return cOceanBlue
	case ColorOceanTwo:
		return cOceanBlue
	case ColorOceanThree:
		return cOceanBlue
	case ColorOceanFour:
		return cOceanBlue
	case ColorPlain:
		return cPlain
	case ColorPlainTwo:
		return cPlain
	case ColorDesert:
		return cDesert
	case ColorDesertTwo:
		return cDesert
	case ColorMountain:
		return cMountain
	case ColorMountainTwo:
		return cMountain
	case ColorSwamp:
		return cSwamp
	case ColorSwampTwo:
		return cSwamp
	case ColorForest:
		return cForest
	case ColorForestTwo:
		return cForest
	case ColorMountainThree:
		return cSixteen
	case ColorForestThree:
		return cOlive
	case ColorOrange:
		return cOrange
	}
	return cAliceBlue
}
