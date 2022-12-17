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

type Terrain int

const (
	TerrNone Terrain = iota
	TerrLand
	TerrOcean
	TerrForest
	TerrSwamp
	TerrMountain
	TerrPlain
	TerrDesert
	TerrWater
	TerrIsland
	TerrStoneCir // circle of stones
	TerrGrove    // mallorn grove
	TerrBog
	TerrCave
	TerrCity
	TerrGuild
	TerrGrave
	TerrRuins
	TerrBattlefield
	TerrEnchFor // enchanted forest
	TerrRockyHill
	TerrTreeCir
	TerrPits
	TerrPasture
	TerrOasis
	TerrYewGrove
	TerrSandPit
	TerrSacGrove // sacred grove
	TerrPopField // poppy field
	TerrTemple
	TerrLair // dragon lair
)
