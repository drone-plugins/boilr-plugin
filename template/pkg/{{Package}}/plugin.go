// Copyright (c) {{ Year }}, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package {{ Package }}

import (
	drone_plugin "github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
)

type (
	// Config for the Plugin.
	Config struct {
		Build    drone_plugin.Build
		Repo     drone_plugin.Repo
		Commit   drone_plugin.Commit
		Stage    drone_plugin.Stage
		Step     drone_plugin.Step
		SemVer   drone_plugin.SemVer
		Settings Settings
	}

	// Plugin interface.
	Plugin interface {
		// Validate the Plugin.
		//
		// If the validation fails then an error should be produced. The Exec
		// method should not be called before Validate.
		Validate() error

		// Exec the Plugin.
		//
		// If the Plugin fails to execute then an error is produced.
		Exec() error
	}

	pluginImpl struct {
		config Config
	}
)

// New Plugin from the given Config.
func New(config Config) Plugin {
	return &pluginImpl{
		config: config,
	}
}
