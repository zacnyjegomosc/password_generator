# Password generator

[![Build Status](https://travis-ci.com/zacnyjegomosc/password_generator.svg?branch=master)](https://travis-ci.com/zacnyjegomosc/password_generator) 
[![Go Report Card](https://goreportcard.com/badge/github.com/zacnyjegomosc/password_generator)](https://goreportcard.com/report/github.com/zacnyjegomosc/password_generator)
[![codecov](https://codecov.io/gh/zacnyjegomosc/password_generator/branch/master/graph/badge.svg)](https://codecov.io/gh/zacnyjegomosc/password_generator)

## Simple and cryptographic secure random password generator (or another string).

This is a small private command line tool, written in Go for learning and having fun (and to have a straightforward, fast and secure
password generator close at hand).

## How to install?

```bash
$ go get -u github.com/zacnyjegomosc/password_generator
```

## Usage

If you want to generate one random string (lenght=32), just type:
```bash
$ password_generator
s2pVJRY9qgpKVFJoHrrqDzIHszFe0D4l
```

Also, you can provide some arguments:
```bash
$ password_generator --help
Usage of password_generator:
  -count int
        How many strings do you need? (default 1)
  -len int
        What length of string do you need? (default 32)

```

So an extended example will be:
```bash
$ password_generator --count 3 --len 12
YWHce8m8uyfk
0nYSsUkMjW7n
yx_8utaH66qq
```

## License
   
This code is licensed under the MIT license.