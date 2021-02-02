// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestRegistration(t *testing.T) {
	testRegistration(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestRegistrationParams(t *testing.T) {
	testRegistrationParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentRegistrationOptions(t *testing.T) {
	testTextDocumentRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestUnregistration(t *testing.T) {
	testUnregistration(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestUnregistrationParams(t *testing.T) {
	testUnregistrationParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
