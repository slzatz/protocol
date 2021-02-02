// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestDidOpenTextDocumentParams(t *testing.T) {
	testDidOpenTextDocumentParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	testDidChangeTextDocumentParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocument(t *testing.T) {
	testTextDocument(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentChangeEvent(t *testing.T) {
	testTextDocumentChangeEvent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentWillSaveEvent(t *testing.T) {
	testTextDocumentWillSaveEvent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	testTextDocumentContentChangeEvent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	testTextDocumentChangeRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	testWillSaveTextDocumentParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	testDidSaveTextDocumentParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	testTextDocumentSaveRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	testDidCloseTextDocumentParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
