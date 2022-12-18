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

type Color int

const (
	ColorNone     Color = iota
	ColorOcean          = 1
	ColorTwo            = 2
	ColorThree          = 3
	ColorFour           = 4
	ColorPlain          = 5
	ColorSix            = 6
	ColorDesert         = 7
	ColorEight          = 8
	ColorMountain       = 9
	ColorTen            = 10
	ColorSwamp          = 11
	ColorTwelve         = 12
	ColorForest         = 13
	ColorFourteen       = 14
	ColorSixteen        = 16
	ColorNineteen       = 19
	ColorNegOne         = -1
)

func (c Color) ToFill() string {
	switch c {
	case ColorOcean:
		return "darkblue"
	case ColorTwo:
		return "cornflowerblue"
	case ColorThree:
		return "lightsteelblue"
	case ColorFour:
		return "blue"
	case ColorPlain:
		return "navajowhite"
	case ColorSix:
		return "brown"
	case ColorDesert:
		return "sandybrown"
	case ColorEight:
		return "khaki"
	case ColorMountain:
		return "sienna"
	case ColorTen:
		return "saddlebrown"
	case ColorSwamp:
		return "yellowgreen"
	case ColorTwelve:
		return "seagreen"
	case ColorForest:
		return "forestgreen"
	case ColorFourteen:
		return "darkolivegreen"
	case ColorSixteen:
		return "peru"
	case ColorNineteen:
		return "olive"
	case ColorNegOne:
		return "orange"
	}
	return "aliceblue"
}

var (
	cAliceBlue, cDesert, cForest, cMountain, cOceanBlue, cOlive, cOrange, cPlain, cSixteen, cSwamp struct {
		r, g, b float64
	}
)

func init() {
	cAliceBlue.r, cAliceBlue.g, cAliceBlue.b = 0xfa, 0xf8, 0xff
	cDesert.r, cDesert.g, cDesert.b = 0xf4, 0xa4, 0x60
	cForest.r, cForest.g, cForest.b = 0x22, 0x8b, 0x22
	cMountain.r, cMountain.g, cMountain.b= 0xa0, 0x52, 0x5d
	cOceanBlue.r,cOceanBlue.g,cOceanBlue.b = 0x1e, 0x90, 0xff
	cOlive.r, cOlive.g, cOlive.b = 0x80, 0x80, 0x00
	cOrange.r, cOrange.g, cOrange.b = 0xff, 0xa5, 0x00
	cPlain.r,cPlain.g,cPlain.b = 0xff, 0xde, 0xad
	cSixteen.r, cSixteen.g, cSixteen.b = 0xcd, 0x85, 0x3f
	cSwamp.r, cSwamp.g, cSwamp.b = 0x9a, 0xcd, 0x32
}
func (c Color) ToRGB() (r,g,b float64) {
	switch c {
	case ColorOcean:
		return cOceanBlue.r,cOceanBlue.g,cOceanBlue.b
	case ColorTwo:
		return cOceanBlue.r,cOceanBlue.g,cOceanBlue.b
	case ColorThree:
		return cOceanBlue.r,cOceanBlue.g,cOceanBlue.b
	case ColorFour:
		return cOceanBlue.r,cOceanBlue.g,cOceanBlue.b
	case ColorPlain:
		return cPlain.r, cPlain.g, cPlain.b
	case ColorSix:
		return cPlain.r, cPlain.g, cPlain.b
	case ColorDesert:
		return cDesert.r, cDesert.g, cDesert.b
	case ColorEight:
		return cDesert.r, cDesert.g, cDesert.b
	case ColorMountain:
		return cMountain.r, cMountain.g, cMountain.b
	case ColorTen:
		return cMountain.r, cMountain.g, cMountain.b
	case ColorSwamp:
		return cSwamp.r, cSwamp.g, cSwamp.b
	case ColorTwelve:
		return cSwamp.r, cSwamp.g, cSwamp.b
	case ColorForest:
		return cForest.r, cForest.g, cForest.b
	case ColorFourteen:
		return cForest.r, cForest.g, cForest.b
	case ColorSixteen:
		return cSixteen.r, cSixteen.g, cSixteen.b
	case ColorNineteen:
		return cOlive.r, cOlive.g, cOlive.b
	case ColorNegOne:
		return cOrange.r, cOrange.g, cOrange.b
	}
	return cAliceBlue.r, cAliceBlue.g, cAliceBlue.b
}
