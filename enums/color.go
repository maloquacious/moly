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
