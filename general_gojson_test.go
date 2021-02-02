// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestWorkspaceFolders(t *testing.T) {
	testWorkspaceFolders(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestInitializeParams(t *testing.T) {
	testInitializeParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestWorkspaceClientCapabilities(t *testing.T) {
	testWorkspaceClientCapabilities(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesSynchronization(t *testing.T) {
	testTextDocumentClientCapabilitiesSynchronization(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesCompletion(t *testing.T) {
	testTextDocumentClientCapabilitiesCompletion(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesHover(t *testing.T) {
	testTextDocumentClientCapabilitiesHover(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesSignatureHelp(t *testing.T) {
	testTextDocumentClientCapabilitiesSignatureHelp(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesReferences(t *testing.T) {
	testTextDocumentClientCapabilitiesReferences(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentHighlight(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentSymbol(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesFormatting(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesRangeFormatting(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesOnTypeFormatting(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesDeclaration(t *testing.T) {
	testTextDocumentClientCapabilitiesDeclaration(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesDefinition(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesTypeDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesTypeDefinition(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesImplementation(t *testing.T) {
	testTextDocumentClientCapabilitiesImplementation(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesCodeAction(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeAction(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesCodeLens(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeLens(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesDocumentLink(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentLink(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesColorProvider(t *testing.T) {
	testTextDocumentClientCapabilitiesColorProvider(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesRename(t *testing.T) {
	testTextDocumentClientCapabilitiesRename(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T) {
	testTextDocumentClientCapabilitiesPublishDiagnostics(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilitiesFoldingRange(t *testing.T) {
	testTextDocumentClientCapabilitiesFoldingRange(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestClientCapabilities(t *testing.T) {
	testClientCapabilities(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestInitializeResult(t *testing.T) {
	testInitializeResult(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	testDocumentLinkRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestInitializedParams(t *testing.T) {
	testInitializedParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
