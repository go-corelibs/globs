[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/globs)
[![codecov](https://codecov.io/gh/go-corelibs/globs/graph/badge.svg?token=tkQUrYOFew)](https://codecov.io/gh/go-corelibs/globs)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/globs)](https://goreportcard.com/report/github.com/go-corelibs/globs)

# globs - glob pattern matching utilities

globs provides simple utilties for working with glob-style pattern matching.

# Installation

``` shell
> go get github.com/go-corelibs/globs@latest
```

# Examples

## Parse, Match, Find

``` go
func main() {
    list, err := Parse("*.go", "*.txt")
    // err == nil
    // list.String() == `["*.go"]`
    if matched, _ := list[0].Match("source.go"); matched {
        // specifically matched with the first glob in the list
    }
    if matched := list.Match("source.go"); matched {
        // one of the globs in the list has matched
    }
    if matched := list.Find("sournce.go"); matched != nil {
        // matched is the *glob.Glob instance that matched the input
    }
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
