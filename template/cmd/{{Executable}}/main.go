// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// DO NOT MODIFY THIS FILE DIRECTLY

package main

import (
	"fmt"
	"os"

	"github.com/drone-plugins/drone-plugin-lib/errors"
	"github.com/drone-plugins/drone-plugin-lib/urfave"
	"github.com/urfave/cli/v2"
	"{{ RepoHost }}/{{ RepoOwner }}/{{ RepoName }}/plugin"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "{{ Executable }}"
	app.Usage = "{{ Usage }}"
	app.Version = version
	app.Action = run
	app.Flags = append(settingsFlags(), urfave.Flags()...)

	if err := app.Run(os.Args); err != nil {
		errors.HandleExit(err)
	}
}

func run(ctx *cli.Context) error {
	urfave.LoggingFromContext(ctx)

	plugin := plugin.New(
		settingsFromContext(ctx),
		urfave.PipelineFromContext(ctx),
		urfave.NetworkFromContext(ctx),
	)

	if err := plugin.Validate(); err != nil {
		return err
	}

	if err := plugin.Execute(); err != nil {
		return err
	}

	return nil
}
