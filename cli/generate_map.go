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

package cli

import (
	"fmt"
	"github.com/playbymail/moly/maps"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// cmdGenerateMap runs the map generator command
var cmdGenerateMap = &cobra.Command{
	Use:   "map",
	Short: "generate a new map",
	RunE: func(cmd *cobra.Command, args []string) error {
		if argsGenerateMap.mapFileName == "" {
			return fmt.Errorf("missing map-data parameter")
		}

		mc, err := maps.Read(argsGenerateMap.mapFileName)
		if err != nil {
			log.Fatal(err)
		}

		if argsGenerateMap.createPNG {
			cobra.CheckErr(mc.ToPNG("worldmap.png"))
		}
		if argsGenerateMap.createSVG {
			svg := mc.ToSvg(true, argsGenerateMap.noColor)
			cobra.CheckErr(os.WriteFile("worldmap.svg", svg.Bytes(!argsGenerateMap.noScale), 0644))
		}

		return nil
	},
}

var argsGenerateMap struct {
	mapFileName       string
	createPNG         bool
	createSVG         bool
	cityFileName      string
	continentFileName string
	gateFileName      string
	landFileName      string
	locationFileName  string
	noColor           bool
	noScale           bool
	regionFileName    string
	roadFileName      string
	seedFileName      string
}
