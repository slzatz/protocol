// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestShowMessageParams(t *testing.T) {
	testShowMessageParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestShowMessageRequestParams(t *testing.T) {
	testShowMessageRequestParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestMessageActionItem(t *testing.T) {
	testMessageActionItem(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestLogMessageParams(t *testing.T) {
	testLogMessageParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
