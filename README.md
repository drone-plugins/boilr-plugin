# boilr-plugin

[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)

This is a [boilr template](https://github.com/tmrts/boilr) for creating a plugin for Drone CI. Get started by installing the template:

```console
$ boilr template download drone-plugins/boilr-plugin drone-plugin -f
```

Create a project in directory my-plugin:

```console
$ boilr template use drone-plugin my-plugin
```

Select if the docker images should be built for Windows:

```text
[?] Please choose a value for "UseWindows" [default: true]:
```

Enter the owner for DockerHub:

```text
[?] Please choose a value for "DockerOwner" [default: "plugins"]:
```

Enter the repository for DockerHub:

```text
[?] Please choose a value for "DockerRepo" [default: ""]: dockerhub
```

Enter the name of the plugin executable:

```text
[?] Please choose a value for "Executable" [default: ""]: drone-dockerhub
```

Enter the git repository host:

```text
[?] Please choose a value for "RepoHost" [default: "github.com"]:
```

Enter the git repository owner:

```text
[?] Please choose a value for "RepoOwner" [default: "drone-plugins"]:
```

Enter the git repository name:

```text
[?] Please choose a value for "RepoName" [default: ""]: drone-dockerhub
```

Enter the usage information of the app:

```text
[?] Please choose a value for "AppUsage" [default: ""]: trigger builds on dockerhub
```

Select if the docker image should use a multiarch base image:

```text
[?] Please choose a value for "UseMultiArch" [default: true]:
```

Enter the maintainer used within the Dockerfile:

```text
[?] Please choose a value for "DockerMaintainer" [default: "Drone.IO Community <drone-dev@googlegroups.com>"]:
```

Enter the schema name used within the Dockerfile:

```text
[?] Please choose a value for "DockerSchemaName" [default: ""]: Drone DockerHub
```

Enter the schema vendor used within the Dockerfile:

```text
[?] Please choose a value for "DockerSchemaVendor" [default: ""]: Drone.IO Community
```

## Example

```text
▶ boilr template use drone-plugin /home/octocat/drone-curseforge -l debug
[✔] Created .dockerignore
[?] Please choose a value for "UseWindows" [default: true]:
[?] Please choose a value for "DockerOwner" [default: "plugins"]:
[?] Please choose a value for "DockerRepo" [default: ""]: curseforge
[?] Please choose a value for "Executable" [default: ""]: drone-curseforge
[✔] Created .drone.star
[✔] Created .gitignore
[✔] Created LICENSE
[?] Please choose a value for "RepoHost" [default: "github.com"]:
[?] Please choose a value for "RepoOwner" [default: "drone-plugins"]:
[?] Please choose a value for "RepoName" [default: ""]: drone-curseforge
[✔] Created cmd/drone-curseforge/config.go
[?] Please choose a value for "Usage" [default: ""]: publish files to curseforge
[✔] Created cmd/drone-curseforge/main.go
[?] Please choose a value for "UseMultiArch" [default: true]:
[?] Please choose a value for "DockerMaintainer" [default: "Drone.IO Community <drone-dev@googlegroups.com>"]:
[?] Please choose a value for "DockerSchemaName" [default: ""]: Drone CurseForge
[?] Please choose a value for "DockerSchemaVendor" [default: "Drone.IO Community"]:
[✔] Created docker/Dockerfile.linux.amd64
[✔] Created docker/Dockerfile.linux.arm
[✔] Created docker/Dockerfile.linux.arm64
[✔] Created docker/Dockerfile.windows.1809
[✔] Created docker/Dockerfile.windows.1903
[✔] Created docker/Dockerfile.windows.1909
[✔] Created docker/manifest.tmpl
[✔] Created go.mod
[✔] Created plugin/impl.go
[✔] Created plugin/impl_test.go
[✔] Created plugin/plugin.go
[✔] Created plugin/plugin_test.go
[✔] Successfully executed the project template drone-plugin in /home/octocat/drone-curseforge
```
