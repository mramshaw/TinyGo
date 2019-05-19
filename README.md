# TinyGo

As a longtime [Golang](http://github.com/mramshaw/Golang) fan, I have cross-compiled
Go programs against [MUSL](#musl) and experimented with
[MQTT](https://github.com/mramshaw/MQTT_and_mosquitto)
(an IoT protocol for sensors and devices such as the Arduino or Raspberry Pi).
I have also experimented with the
[Raspberry Pi](https://github.com/mramshaw/Speech-Recognition#raspberry-pi)
(a single-board microcomputer) so TinyGo sounded like it was worth a look.

## What is TinyGo?

> TinyGo is a Go compiler intended for use in small places such as microcontrollers,
> WebAssembly (WASM), and command-line tools.

And:

> TinyGo is a project to bring Go to microcontrollers and small systems with a single processor core.

Both of the above quotes are from:

    http://github.com/tinygo-org/tinygo

## Prerequisites

Requires __Go v1.11__ or greater.

Verify as follows:

```bash
$ go version
go version go1.11 linux/amd64
$
```

Requires libstdc++6 __3.4.22__ or greater. Verify this as follows:

```bash
$ strings /usr/lib/x86_64-linux-gnu/libstdc++.so.6 | grep GLIBCXX
```

If `GLIBCXX_3.4.22` is not listed, then libstdc++6 must be installed or upgraded.

On Ubuntu, this can be done as follows:

```bash
$ sudo add-apt-repository ppa:ubuntu-toolchain-r/test
...
$ sudo apt-get update
...
$ sudo apt-get upgrade libstdc++6
...
$
```

## Installation

As per the instructions, download and install as follows:

```bash
$ wget https://github.com/tinygo-org/tinygo/releases/download/v0.5.0/tinygo_0.5.0_amd64.deb
...
$ sudo dpkg -i tinygo_0.5.0_amd64.deb
...
$
```

Note that the version number is subject to change. Check for the latest release here:

    http://github.com/tinygo-org/tinygo/releases

Verify the installation as follows:

```bash
$ /usr/local/tinygo/bin/tinygo version
tinygo version 0.5.0 linux/amd64
$
```

[By default TinyGo is installed in __/usr/local/tinygo/__ with __root__ permissions but globally executable.]

## Targets

Straight out of the box, TinyGo can be used to compile to [WebAssembly](#webassembly).

For hardware platforms such as ARM or AVR there are additional requirements.

The full list of supported boards is provided here:

    http://github.com/tinygo-org/tinygo#supported-boardstargets

#### ARM

ARM processors are apparently well supported.

From:

    https://tinygo.org/compiler-internals/microcontrollers/

> ARM Cortex-M processors are well supported.

ARM devices apparently require __clang-8__.

#### AVR

Arduino (AVR) devices apparently require __gcc-avr__, __avr-libc__ and __avrdude__.

## Running

We will try it out for WASM as follows:

#### Compile WASM

```bash
$ cd src
$ /usr/local/tinygo/bin/tinygo build -o ./wasm.wasm -target wasm ./wasm.go
$
```

#### Run webserver

Launch a web server to serve up our WASM-enriched HTML:

```bash
$ cd ..
$ go run server.go
2019/05/18 16:42:41 Serving './src' on http://localhost:8080
^Csignal: interrupt
$
```

This is important as our WASM needs to be properly MIME-encoded.

As usual, <kbd>Ctrl-C</kbd> to terminate once testing is complete.

#### Open browser

Open a javascript-enabled browser to the following URL:

    http://localhost:8080/

And test. The results should look as follows:

![WASM working](images/WASM_working.png)

## Reference

Some useful references are listed below.

#### TinyGo references

The TinyGo repo can be found here:

    http://github.com/tinygo-org/tinygo

Getting started with TinyGo:

    http://tinygo.org/getting-started/linux/

#### LLVM

Probably the definitive site for learning about LLVM:

    http://llvm.org/

#### MUSL

Probably the definitive site for learning about MUSL:

    http://www.musl-libc.org/

MUSL is used as the standard C library by the Alpine Linux distribution.

According to Wikipedia:

> Musl was designed from scratch to allow efficient static linking and to have realtime-quality robustness
> by avoiding races, internal failures on resource exhaustion and various other bad worst-case behaviors
> present in existing implementations

And:

>It also implements most of the widely used non-standard Linux, BSD, and glibc functions.

Both of the above quotes are from:

    http://en.wikipedia.org/wiki/Musl

The second quote is particularly important; MUSL only implements ***most*** of the standard
glibc functions, which means it cannot be simply used as a drop-in replacement for glibc.

Nevertheless, if used carefully it can be used as a replacement for ***most*** uses.

For more information, please refer to the following link:

    http://wiki.musl-libc.org/compatibility.html

Likewise, I expect TinyGo to be merely a ***partial*** implementation of Go, as is usely the
case with micro-languages.

#### WebAssembly

> WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine.
> Wasm is designed as a portable target for compilation of high-level languages like C/C++/Rust,
> enabling deployment on the web for client and server applications.

    http://webassembly.org

WebAssembly seems to be primarily associated with Rust but has been a topic of interest in the
Golang community for some time. For instance:

    http://changelog.com/gotime/bonus-77

[At various Go meetups I have attended, WASM has been discussed.]

## Credits

Inspired by this podcast:

    http://changelog.com/gotime/84

## To Do

- [ ] More testing
