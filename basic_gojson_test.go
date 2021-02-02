// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestPosition(t *testing.T) { testPosition(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestRange(t *testing.T) { testRange(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestLocation(t *testing.T) { testLocation(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestLocationLink(t *testing.T) {
	testLocationLink(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDiagnostic(t *testing.T) { testDiagnostic(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestDiagnosticRelatedInformation(t *testing.T) {
	testDiagnosticRelatedInformation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCommand(t *testing.T) { testCommand(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestTextEdit(t *testing.T) { testTextEdit(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestTextDocumentEdit(t *testing.T) {
	testTextDocumentEdit(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCreateFileOptions(t *testing.T) {
	testCreateFileOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCreateFile(t *testing.T) { testCreateFile(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestRenameFileOptions(t *testing.T) {
	testRenameFileOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestRenameFile(t *testing.T) { testRenameFile(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestDeleteFileOptions(t *testing.T) {
	testDeleteFileOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDeleteFile(t *testing.T) { testDeleteFile(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestWorkspaceEdit(t *testing.T) {
	testWorkspaceEdit(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentIdentifier(t *testing.T) {
	testTextDocumentIdentifier(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentItem(t *testing.T) {
	testTextDocumentItem(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestVersionedTextDocumentIdentifier(t *testing.T) {
	testVersionedTextDocumentIdentifier(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentPositionParams(t *testing.T) {
	testTextDocumentPositionParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentFilter(t *testing.T) {
	testDocumentFilter(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentSelector(t *testing.T) {
	testDocumentSelector(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestMarkupContent(t *testing.T) {
	testMarkupContent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
