// Copyright (c) {{ Year }}, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"github.com/urfave/cli"

	"{{ Host }}/{{ Owner }}/{{ Repo }}/pkg/{{ Package }}"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags() []cli.Flag {
	// Replace below with all the flags required for the plugin's specific
	// settings.
	return []cli.Flag{}
}

// settingsFromContext creates a plugin.Settings from the cli.Context.
func settingsFromContext(ctx *cli.Context) {{ Package }}.Settings {
	// Replace below with the parsing of the
	return plugin.Settings{}
}

// modifyFlags allows modification of the cli.Flags before the application is
// run.
func modifyFlags(flags []cli.Flag) []cli.Flag {
	// Replace below with any changes to the flags required by the application.
	//
	// One use for this is showing flags from Drone's environment that may be
	// useful during the development of the plugin.
	return flags
}
