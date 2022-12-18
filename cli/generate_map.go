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
	"github.com/maloquacious/moly/maps"
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
		svg := mc.ToSvg(true, argsGenerateMap.noColor)
		cobra.CheckErr(os.WriteFile("worldmap.svg", svg.Bytes(!argsGenerateMap.noScale), 0644))
		cobra.CheckErr(mc.ToPNG("worldmap.png"))

		return nil
	},
}

var argsGenerateMap struct {
	mapFileName       string
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

func init() {
	cmdGenerate.AddCommand(cmdGenerateMap)
	// inputs
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.mapFileName, "map-data", "worldmap.txt", "map data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.cityFileName, "city-data", "cities.json", "city name data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.landFileName, "land-data", "lands.json", "land data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.regionFileName, "region-data", "regions.json", "region data to import")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.noColor, "no-color", argsGenerateMap.noColor, "do not color map")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.noScale, "no-scale", argsGenerateMap.noScale, "do not scale map")

	// outputs
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.continentFileName, "continent-data", "continents.json", "continent data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.locationFileName, "location-data", "locations.json", "location data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.gateFileName, "gate-data", "gates.json", "gate data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.roadFileName, "road-data", "roads.json", "road data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.seedFileName, "seed-data", "randseed.json", "random seed data to export")
}
