// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestWorkspaceFolder(t *testing.T) {
	testWorkspaceFolder(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidChangeWorkspaceFoldersParams(t *testing.T) {
	testDidChangeWorkspaceFoldersParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestWorkspaceFoldersChangeEvent(t *testing.T) {
	testWorkspaceFoldersChangeEvent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidChangeConfigurationParams(t *testing.T) {
	testDidChangeConfigurationParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestConfigurationParams(t *testing.T) {
	testConfigurationParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestConfigurationItem(t *testing.T) {
	testConfigurationItem(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidChangeWatchedFilesParams(t *testing.T) {
	testDidChangeWatchedFilesParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestFileEvent(t *testing.T) {
	testFileEvent(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestDidChangeWatchedFilesRegistrationOptions(t *testing.T) {
	testDidChangeWatchedFilesRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestWorkspaceSymbolParams(t *testing.T) {
	testWorkspaceSymbolParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestExecuteCommandParams(t *testing.T) {
	testExecuteCommandParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestExecuteCommandRegistrationOptions(t *testing.T) {
	testExecuteCommandRegistrationOptions(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestApplyWorkspaceEditParams(t *testing.T) {
	testApplyWorkspaceEditParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}

func TestApplyWorkspaceEditResponse(t *testing.T) {
	testApplyWorkspaceEditResponse(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
