// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

// OptionalTrue return reference of true for a optional bool
func OptionalTrue() *bool {
	v := true
	return &v
}

// OptionalFalse return reference of false for a optional bool
func OptionalFalse() *bool {
	v := false
	return &v
}
