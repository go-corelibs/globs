// Copyright (c) 2024  The Go-Curses Authors
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

package globs

import (
	"testing"

	glob "github.com/ganbarodigital/go_glob"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGlobs(t *testing.T) {
	Convey("Parse", t, func() {
		list, err := Parse()
		So(err, ShouldEqual, nil)
		So(list, ShouldEqual, Globs(nil))
		list, err = Parse(
			"*.txt",
			".[a-z]*",
		)
		So(err, ShouldEqual, nil)
		So(Globs(list).String(), ShouldEqual, `["*.txt",".[a-z]*"]`)
		list, err = Parse("[nope")
		So(err, ShouldNotEqual, nil)
		So(list, ShouldEqual, Globs(nil))
	})

	Convey("Match", t, func() {
		list, err := Parse("*.txt", "*.go")
		So(err, ShouldEqual, nil)
		So(len(list), ShouldEqual, 2)
		So(list.Match("nope.log"), ShouldEqual, false)
		So(list.Match("source.go"), ShouldEqual, true)
	})

	Convey("Find", t, func() {
		list, err := Parse("*.txt", "*.go")
		So(err, ShouldEqual, nil)
		So(len(list), ShouldEqual, 2)
		So(list.Find("nope.log"), ShouldEqual, (*glob.Glob)(nil))
		So(list.Find("source.go"), ShouldEqual, list[1])
	})
}
