# mdviewer

[![PkgGoDev](https://pkg.go.dev/badge/github.com/noborus/mdviewer?lightbox=false)](https://pkg.go.dev/github.com/noborus/mdviewer)
[![Latest Release](https://img.shields.io/github/v/release/noborus/mdviewer.svg?lightbox=false)](https://github.com/noborus/mdviewer/releases)

This is a sample using the [ov(oviewer)](https://github.com/noborus/ov) package.

![docs/mdviewer.gif](docs/mdviewer.gif)

Start with `mdviewer markdown file`.

```sh
mdviewer README.md
```

Rendered and displayed by [glamour](https://github.com/charmbracelet/glamour).

* Press the `]` key to see the original file.
* Press the `[` key to return.

See the [ov](https://github.com/noborus/ov) for other operations.

## Installation

### go install

```sh
go install github.com/noborus/mdviewer@latest
```

### Binary

[Releases](https://github.com/noborus/mdviewer/releases/latest)

### Homebrew

```sh
brew install noborus/tap/mdviewer
```

### Arch Linux

[https://aur.archlinux.org/packages/mdviewer-git](https://aur.archlinux.org/packages/mdviewer-git)

### deb

You can download the package from [releases](https://github.com/noborus/mdviewer/releases).

```console
curl -L -O https://github.com/noborus/mdviewer/releases/download/vx.x.x/mdviewer_x.x.x-1_amd64.deb
sudo apt install ./mdviewer_x.x.x_amd64.deb
```

### rpm

You can download the package from [releases](https://github.com/noborus/mdviewer/releases).

```console
sudo rpm -ivh https://github.com/noborus/mdviewer/releases/download/vx.x.x/mdviewer_x.x.x-1_amd64.rpm
```
