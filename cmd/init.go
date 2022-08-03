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
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Package struct {
	Name    string   `toml:"name"`
	Version string   `toml:"version"`
	Authors []string `toml:"authors"`
}

type Manifest struct {
	Package      Package           `toml:"package"`
	Dependencies map[string]string `toml:"dependencies"`
}

const main = `fn main()
  println("Hello, world!)
end
`

const ignore = `/target`

// initCmd represents the init command
var (
	version string
	authors []string
	initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "Initalize a package with an optional name",
		Long:  `TODO`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			manifest := Manifest{
				Package: Package{
					Name:    "",
					Version: version,
					Authors: authors,
				},
				Dependencies: make(map[string]string),
			}
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			wds := strings.Split(wd, "/")
			manifest.Package.Name = wds[len(wds)-1]
			if len(args) == 1 {
				manifest.Package.Name = args[0]
			}
			f, err := os.Create("./graphite.toml")
			if err != nil {
				return err
			}
			b, err := toml.Marshal(manifest)
			if err != nil {
				return err
			}
			f.Write(b)

			err = os.Mkdir("src", fs.ModePerm)
			if err != nil {
				return err
			}
			s, err := os.Create("./src/main.graphene")
			if err != nil {
				return err
			}
			s.Write([]byte(main))

			i, err := os.Create("./.gitignore")
			if err != nil {
				return err
			}
			i.Write([]byte(ignore))

			git := exec.Command("git", "init")

			err = git.Run()

			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVarP(&version, "version", "v", "0.1.0", "the inital version of the app")
	initCmd.PersistentFlags().StringArrayVarP(&authors, "authors", "a", []string{}, "a list of project authors")
}
