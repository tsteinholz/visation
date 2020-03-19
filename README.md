![Visation Header](https://github.com/tsteinholz/visation/raw/master/.github/visation-header2.png)
A graphical application to provide a visual sensation to your music!

![GitHub](https://img.shields.io/github/license/tsteinholz/visation?style=for-the-badge) ![Travis (.com)](https://img.shields.io/travis/com/tsteinholz/visation?style=for-the-badge) ![Code Report](https://goreportcard.com/badge/github.com/tsteinholz/visation?style=for-the-badge) ![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/tsteinholz/visation?style=for-the-badge)

## Building

1. [Install Go](https://golang.org/doc/install)
2. [Install Vulkan SDK](https://vulkan.lunarg.com/doc/view/1.1.106.0/linux/getting_started.html)
2. Clone & Build...

```bash
go get github.com/tsteinholz/visation
cd $GOPATH/src/github.com/tsteinholz/visation/desktop
git submodule init
git submodule upgrade
go build
./visation
```
