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

import "fmt"

// The entity number of a region determines where it is on the map.
// Here is how:
//
//   (r,c)
// 	+-------------------+
// 	|(1,1)        (1,99)|
// 	|                   |
// 	|                   |
// 	|(n,1)        (n,99)|
// 	+-------------------+
//
// Entity [10101] corresponds to (1, 1).
// Entity [10199] corresponds to (1,99).
//
// Note that for player convenience an alternate method of representing
// location entity numbers may be used, i.e. 'aa'. 101, 'ab' . 102,
// so [aa01] . [10101], [ab53] . [10253].

func rowColToRegion(row, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}
