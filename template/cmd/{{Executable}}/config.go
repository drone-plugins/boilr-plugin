// Copyright (c) 2021, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"github.com/urfave/cli/v2"
	"{{ RepoHost }}/{{ RepoOwner }}/{{ RepoName }}/plugin"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	// Replace below with all the flags required for the plugin.
	// Use Destination within the cli.Flags to populate settings
	return []cli.Flag{}
}
