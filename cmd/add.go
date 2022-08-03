/*
Copyright © 2022 Devin Rockwell

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dependency",
	Long:  `TODO`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var manifest Manifest = Manifest{}
		v, err := ioutil.ReadFile("./graphite.toml")
		if err != nil {
			return err
		}
		err = toml.Unmarshal(v, &manifest)
		if err != nil {
			return err
		}
		nv := strings.Split(args[0], "@")
		if len(nv) != 2 {
			return errors.New("incorrect format must be NAME@VERSION")
		}
		if manifest.Dependencies == nil {
			manifest.Dependencies = make(map[string]string)
		}
		manifest.Dependencies[nv[0]] = nv[1]
		f, err := os.Create("./graphite.toml")
		if err != nil {
			return err
		}
		b, err := toml.Marshal(manifest)
		if err != nil {
			return err
		}
		f.Write(b)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
