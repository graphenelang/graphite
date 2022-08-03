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

/*
Graphite is a package manager for graphene


Usage:
  graphite [command]

Available Commands:
  add         Add a dependency
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initalize a package with an optional name
  rm          Remove dependency

Flags:
  -h, --help     help for graphite
*/
package main

import "github.com/graphenelang/graphite/cmd"

func main() {
	cmd.Execute()
}
