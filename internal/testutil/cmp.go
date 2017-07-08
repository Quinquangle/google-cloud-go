// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"math"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

var (
	alwaysEqual = cmp.Comparer(func(_, _ interface{}) bool { return true })

	defaultCmpOptions = []cmp.Option{
		// Use proto.Equal for protobufs
		cmp.Comparer(proto.Equal),
		// NaNs compare equal
		cmp.FilterValues(func(x, y float64) bool {
			return math.IsNaN(x) && math.IsNaN(y)
		}, alwaysEqual),
		cmp.FilterValues(func(x, y float32) bool {
			return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
		}, alwaysEqual),
	}
)

// Equal tests two values for equality.
func Equal(x, y interface{}, opts ...cmp.Option) bool {
	// Put default options at the end. Order doesn't matter.
	opts = append(opts[:len(opts):len(opts)], defaultCmpOptions...)
	return cmp.Equal(x, y, opts...)
}
