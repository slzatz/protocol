// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestCompletionParams(t *testing.T) {
	testCompletionParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCompletionContext(t *testing.T) {
	testCompletionContext(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCompletionList(t *testing.T) {
	testCompletionList(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCompletionItem(t *testing.T) {
	testCompletionItem(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCompletionRegistrationOptions(t *testing.T) {
	testCompletionRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestHover(t *testing.T) { testHover(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestSignatureHelp(t *testing.T) {
	testSignatureHelp(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestSignatureInformation(t *testing.T) {
	testSignatureInformation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestParameterInformation(t *testing.T) {
	testParameterInformation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestSignatureHelpRegistrationOptions(t *testing.T) {
	testSignatureHelpRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestReferenceParams(t *testing.T) {
	testReferenceParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestReferenceContext(t *testing.T) {
	testReferenceContext(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentHighlight(t *testing.T) {
	testDocumentHighlight(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentSymbolParams(t *testing.T) {
	testDocumentSymbolParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentSymbol(t *testing.T) {
	testDocumentSymbol(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCodeActionParams(t *testing.T) {
	testCodeActionParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCodeActionContext(t *testing.T) {
	testCodeActionContext(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCodeAction(t *testing.T) { testCodeAction(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestCodeActionRegistrationOptions(t *testing.T) {
	testCodeActionRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCodeLensParams(t *testing.T) {
	testCodeLensParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestCodeLens(t *testing.T) { testCodeLens(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestCodeLensRegistrationOptions(t *testing.T) {
	testCodeLensRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentLinkParams(t *testing.T) {
	testDocumentLinkParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentLink(t *testing.T) {
	testDocumentLink(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentColorParams(t *testing.T) {
	testDocumentColorParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestColorInformation(t *testing.T) {
	testColorInformation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestColor(t *testing.T) { testColor(t, json.MarshalNoEscape, json.UnmarshalNoEscape) }

func TestColorPresentationParams(t *testing.T) {
	testColorPresentationParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestColorPresentation(t *testing.T) {
	testColorPresentation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentFormattingParams(t *testing.T) {
	testDocumentFormattingParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestSymbolInformation(t *testing.T) {
	testSymbolInformation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestFormattingOptions(t *testing.T) {
	testFormattingOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentRangeFormattingParams(t *testing.T) {
	testDocumentRangeFormattingParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentOnTypeFormattingParams(t *testing.T) {
	testDocumentOnTypeFormattingParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentOnTypeFormattingRegistrationOptions(t *testing.T) {
	testDocumentOnTypeFormattingRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestRenameParams(t *testing.T) {
	testRenameParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestRenameRegistrationOptions(t *testing.T) {
	testRenameRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestFoldingRangeParams(t *testing.T) {
	testFoldingRangeParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestFoldingRange(t *testing.T) {
	testFoldingRange(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
