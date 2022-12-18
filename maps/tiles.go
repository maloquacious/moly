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
	"github.com/maloquacious/moly/enums"
	"math/rand"
	"unicode"
)

type Tiles map[string]*Tile

type Tile struct {
	Row, Col int // map tile we're inside

	City             int
	Depth            int
	Gates            []*Gate
	Inside           int
	Mark             int
	Name             string
	Region           string
	Roads            []*Road
	SaveChar         rune
	SummerbridgeFlag int
	Subs             []int
	UldimFlag        int
	Is               struct {
		Hidden         bool
		MajorCity      bool
		RegionBoundary bool
		SafeHaven      bool
		SeaLane        bool
		Water          bool
	}

	Color   enums.Color // map coloring for output
	Terrain enums.Terrain
}

func toTiles(cells [][]rune) Tiles {
	tiles := make(map[string]*Tile)
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			if cells[row][col] == '#' { // hole in map
				continue
			}

			coords := fmt.Sprintf("%d,%d", row, col)
			tile := &Tile{
				Row:      row,
				Col:      col,
				Region:   rowColToRegion(row, col),
				Depth:    2,
				SaveChar: cells[row][col],
				Color:    enums.ColorOrange,
			}
			tiles[coords] = tile

			switch cells[row][col] {
			case ';':
				tile.Is.SeaLane = true
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOcean
			case ',':
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOcean
			case ':':
				tile.Is.SeaLane = true
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanTwo
			case '.':
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanTwo
			case '~':
				tile.Is.SeaLane = true
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanThree
			case ' ':
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanThree
			case '"':
				tile.Is.SeaLane = true
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanFour
			case '\'':
				tile.Terrain = enums.TerrOcean
				tile.Color = enums.ColorOceanFour
			case 'p':
				tile.Terrain = enums.TerrPlain
				tile.Color = enums.ColorPlain
			case 'P':
				tile.Terrain = enums.TerrPlain
				tile.Color = enums.ColorPlainTwo
			case 'd':
				tile.Terrain = enums.TerrDesert
				tile.Color = enums.ColorDesert
			case 'D':
				tile.Terrain = enums.TerrDesert
				tile.Color = enums.ColorDesertTwo
			case 'm':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountain
			case 'M':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountainTwo
			case 's':
				tile.Terrain = enums.TerrSwamp
				tile.Color = enums.ColorSwamp
			case 'S':
				tile.Terrain = enums.TerrSwamp
				tile.Color = enums.ColorSwampTwo
			case 'f':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForest
			case 'F':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestTwo
			case 'o':
				switch rand.Intn(10) + 1 {
				case 1, 2, 3:
					tile.Terrain = enums.TerrForest
					tile.Color = enums.ColorForest
				case 4, 5, 6:
					tile.Terrain = enums.TerrPlain
					tile.Color = enums.ColorPlain
				case 7, 8:
					tile.Terrain = enums.TerrMountain
					tile.Color = enums.ColorMountainTwo
				case 9:
					tile.Terrain = enums.TerrSwamp
					tile.Color = enums.ColorSwampTwo
				case 10:
					tile.Terrain = enums.TerrDesert
					tile.Color = enums.ColorDesertTwo
				}
			case '?':
				tile.Is.Hidden = true
				tile.Terrain = enums.TerrLand

			// special stuff
			case '^':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountain // was 15, unique
				tile.UldimFlag = 1
				tile.Is.RegionBoundary = true
			case 'v':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountain // was 15, unique
				tile.UldimFlag = 2
				tile.Is.RegionBoundary = true
			case '{':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountainThree
				tile.UldimFlag = 3
				tile.Name = "Uldim pass"
				tile.Is.RegionBoundary = true
			case '}':
				tile.Terrain = enums.TerrMountain
				tile.Color = enums.ColorMountainThree
				tile.UldimFlag = 4
				tile.Name = "Uldim pass"
				tile.Is.RegionBoundary = true
			case ']':
				tile.Terrain = enums.TerrSwamp
				tile.SummerbridgeFlag = 1
				tile.Name = "Summerbridge"
				tile.Is.RegionBoundary = true
			case '[':
				tile.Terrain = enums.TerrSwamp
				tile.SummerbridgeFlag = 2
				tile.Name = "Summerbridge"
				tile.Is.RegionBoundary = true
			case 'O':
				tile.Terrain = enums.TerrMountain
				tile.Color = -1
				tile.Name = "Mt. Olympus"
			case '1':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '2':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '3':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '4':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '5':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '6':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '7':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
				tile.Is.SafeHaven = true
			case '8':
				tile.Terrain = enums.TerrForest
				tile.Color = enums.ColorForestThree
			case '*':
				tile.Terrain = enums.TerrLand
			case '%':
				tile.Terrain = enums.TerrLand
			default:
				if unicode.IsDigit(cells[row][col]) {
					panic(fmt.Sprintf("%d: %d: terrain %q: should not fall through", row+1, col+1, cells[row][col]))
				}
				panic(fmt.Sprintf("%d: %d: terrain %q: unknown", row+1, col+1, cells[row][col]))
			}

			tile.Is.Water = tile.Terrain == enums.TerrWater || tile.Terrain == enums.TerrOcean
		}
	}

	return tiles
}
