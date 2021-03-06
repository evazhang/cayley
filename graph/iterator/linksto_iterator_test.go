// Copyright 2014 The Cayley Authors. All rights reserved.
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

package iterator

import (
	"testing"

	"github.com/google/cayley/quad"
)

func TestLinksTo(t *testing.T) {
	qs := &store{
		data: []string{1: "cool"},
		iter: newFixed(),
	}
	qs.iter.(*Fixed).Add(2)
	fixed := newFixed()
	val := qs.ValueOf("cool")
	if val != 1 {
		t.Fatalf("Failed to return correct value, got:%v expect:1", val)
	}
	fixed.Add(val)
	lto := NewLinksTo(qs, fixed, quad.Object)
	if !lto.Next() {
		t.Error("At least one quad matches the fixed object")
	}
	val = lto.Result()
	if val != 2 {
		t.Errorf("Quad index 2, such as %s, should match %s", qs.Quad(2), qs.Quad(val))
	}
}
