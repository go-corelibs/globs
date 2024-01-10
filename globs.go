// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package globs provides glob pattern matching utilities
package globs

import (
	"fmt"

	glob "github.com/ganbarodigital/go_glob"
)

// Globs is a slice of compiled glob patterns
type Globs []*glob.Glob

// Parse creates and tests new glob instances for each pattern given, stopping
// and returning upon first error encountered
func Parse(patterns ...string) (globs Globs, err error) {
	for _, pattern := range patterns {
		g := glob.NewGlob(pattern)
		if _, err = g.Match("./test/file.txt"); err != nil {
			err = fmt.Errorf("%q error: %w", pattern, err)
			return
		}
		globs = append(globs, g)
	}
	return
}

// String implements the fmt.Stringer interface, outputting the list of globs
// in a JSON string array format
func (s Globs) String() (list string) {
	list += "["
	for idx, g := range s {
		if idx > 0 {
			list += ","
		}
		list += fmt.Sprintf("%q", g.Pattern())
	}
	list += "]"
	return
}

// Match returns true if one of the globs matches the given input
func (s Globs) Match(input string) (matched bool) {
	for _, g := range s {
		if matched, _ = g.Match(input); matched {
			return
		}
	}
	return
}

// Find returns the first glob that matches the given input
func (s Globs) Find(input string) (matched *glob.Glob) {
	for _, g := range s {
		if match, _ := g.Match(input); match {
			matched = g
			return
		}
	}
	return
}
