![Visation Header](https://github.com/tsteinholz/visation/raw/master/.github/visation-header2.png)
A graphical application to provide a visual sensation to your music!

## Building

![](https://api.travis-ci.org/tsteinholz/visation.svg?branch=master)
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
