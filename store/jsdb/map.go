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

package jsdb

import "encoding/json"

// Map contains the data for every hex in the world map
type Map map[string]*Hex

// Hex contains the data for a single hex in the world map
type Hex struct {
	Coords  Coords `json:"-"`
	Terrain string `json:"terrain"`
	City    *City  `json:"city,omitempty"`
	Range   *Range `json:"range,omitempty"`
}

// City contains the data for a city in a single hex
type City struct {
	Name string `json:"name"`
	Hex  *Hex   `json:"-"`
}

// Range contains the data for a range of mountain hexes
type Range struct {
	Name  string `json:"name"`
	Hexes []*Hex `json:"-"`
}

func MapLoad(data []byte) (Map, error) {
	m := make(map[string]*Hex)
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func MapSave(m Map) ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}