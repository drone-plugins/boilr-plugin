// Copyright (c) {{ Year }}, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// DO NOT MODIFY THIS FILE DIRECTLY

package main

import (
	"os"

	"github.com/drone-plugins/drone-plugin-lib/pkg/urfave"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"{{ Host }}/{{ Owner }}/{{ Repo }}/pkg/{{ Package }}"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "nuget plugin"
	app.Usage = "nuget plugin"
	app.Action = run

	// Create the flags for the application
	flags := settingsFlags()
	flags = append(flags, urfave.BuildFlags()...)
	flags = append(flags, urfave.RepoFlags()...)
	flags = append(flags, urfave.CommitFlags()...)
	flags = append(flags, urfave.StageFlags()...)
	flags = append(flags, urfave.StepFlags()...)
	flags = append(flags, urfave.SemVerFlags()...)

	app.Flags = modifyFlags(flags)

	// Run the application
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	config := {{ Package }}.Config{
		Build:    urfave.BuildFromContext(c),
		Repo:     urfave.RepoFromContext(c),
		Commit:   urfave.CommitFromContext(c),
		Stage:    urfave.StageFromContext(c),
		Step:     urfave.StepFromContext(c),
		SemVer:   urfave.SemVerFromContext(c),
		Settings: settingsFromContext(c),
	}

	plugin := {{ Package }}.New(config)

	// Validate the settings
	if err := plugin.Validate(); err != nil {
		return errors.Wrap(err, "validation failed")
	}

	// Run the plugin
	if err := plugin.Exec(); err != nil {
		return errors.Wrap(err, "exec failed")
	}

	return nil
}
