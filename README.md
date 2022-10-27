<div align="center">

<img alt="gif-header" src="https://github.com/lpmatos/personal-resume/blob/main/assets/coding.gif" width="225"/>

<h2>✨ Prune container images in a CLI way ✨</h2>

[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)]()
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)]()
[![GitHub repo size](https://img.shields.io/github/repo-size/lpmatos/ghcr-prune)](https://github.com/lpmatos/ghcr-prune)

---

<img alt="gif-about" src="https://github.com/lpmatos/personal-resume/blob/main/assets/hey.gif" width="300"/>

<p>A CLI tool that prune old images on GitHub (ghcr.io) registry and GitLab (registry.gitlab.com) registry</p>

<p>
  <a href="#getting-started">Getting Started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#usage">Usage</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#installation">Installation</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#concepts">Concepts</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#versioning">Versioning</a>
</p>

</div>

---

## ➤ Getting Started <a name = "getting-started"></a>

If you want contribute on this project, first you need to make a **git clone**:

```bash
git clone --depth 1 https://github.com/ci-monk/drprune.git -b main
```

This will give you access to the code on your **local machine**.

## ➤ Usage <a name = "usage"></a>

**Variables**

| Environment  	| Description                   	|
|--------------	|-------------------------------	|
| GH_TOKEN     	| GitHub API Token              	|
| GH_USERNAME  	| GitHub User/Organization Name 	|
| GH_CONTAINER 	| GitHub Container Name         	|
| GL_TOKEN     	| GitLab API Token              	|
| GL_NAMESPACE 	| GitLab Namespace              	|

### CLI

**GitHub**

```bash
drprune gh images -t $GH_TOKEN -n <username> -c <container-name>
drprune gh insights -t $GH_TOKEN -n <username>
```

**GitLab**

```bash
drprune gl images -t $GL_TOKEN -ns <namespace>
drprune gl insights -t $GL_TOKEN -ns <namespace>
```

### CI

**GitHub Actions**

```yaml

```

**GitLab CI**

```yaml

```

## ➤ Installation <a name = "installation"></a>

with `go`:

```bash
# if you cannot install directly, try following command,
# then input install command again
go get -u github.com/lpmatos/drprune/cmd/drprune

# or
go get -v ./...
go run ./cmd/drprune/main.go --help
```

with `brew`:

```bash
brew tap ci-monk/tools
brew install drprune
```

## ➤ Concepts <a name = "concepts"></a>

### Cobra

Cobra is a CLI framework for Golang. Using it you can speed up your development and creating a powerful and modern CLI application. Cobra is built on a structure of commands, arguments and flags:

- Commands represent actions.
- Args are things.
- Flags are modifiers for those actions.

The best applications will read like sentences when used. Users will know how to use the application because they will natively understand how to use it. This pattern is: `APPCLI VERB NOUN --ADJECTIVE` or `APPCLI COMMAND ARG --FLAG`.

#### Commands

Command is the central point of the application. Each interaction thar the application supports will be contained in a command. We can create commands with children commands and optionally run an `action`. In the example above, `server` is the command.

#### Flags

A flag is a way to modify the behavior of a command. Cobra supports fully POSIX-compliant flags as well the Go flag package. A Cobra command can define flags that persist through to children commands and flags that are only available to that command. In the example above, `port` is the flag.

## ➤ Learnings <a name = "learnings"></a>

- Create a Golang CLI application.
- Create a multistage Golang Dockerfile.
- Create a docker-compose file with waiting entrypoint.
- Setup a Golang Releaser publish pipeline using github actions.
- Understand how GiHub works with packages.
- Understand how GitLab works with packages.
- Handler operations in GitHub API to delete container images using a SDK.
- Handler operations in GitLab API to delete container images using a SDK.

## ➤ Links <a name = "links"></a>

- https://github.com/pterm/pterm
- https://github.com/jedib0t/go-pretty
- https://github.com/charmbracelet/glamour
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/bubbletea

## ➤ Versioning <a name = "versioning"></a>

To check the change history, please access the [**CHANGELOG.md**](CHANGELOG.md) file.

## ➤ Project status <a name = "project-status"></a>

Currently the project is constantly being updated! 👾

## ➤ Donations <a name = "donations"></a>

If my work has impacted your life in a positive way and you'd like to buy me a coffee (or a hundred), that'd be much appreciated!

## ➤ Show your support <a name = "show-your-support"></a>

<div align="center">

Give me a ⭐️ if this project helped you!

<img alt="gif-header" src="https://www.icegif.com/wp-content/uploads/baby-yoda-bye-bye-icegif.gif" width="225"/>

Made with 💜 by [me](https://github.com/lpmatos) 👋 inspired on [readme-md-generator](https://github.com/kefranabg/readme-md-generator)

</div>
