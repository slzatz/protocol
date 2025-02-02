// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

const (
	keyActions                           = "actions"
	keyActiveParameter                   = "activeParameter"
	keyActiveSignature                   = "activeSignature"
	keyActiveSignatureHelp               = "activeSignatureHelp"
	keyAdded                             = "added"
	keyAdditionalTextEdits               = "additionalTextEdits"
	keyAlpha                             = "alpha"
	keyApplied                           = "applied"
	keyApplyEdit                         = "applyEdit"
	keyArguments                         = "arguments"
	keyBlue                              = "blue"
	keyCancellable                       = "cancellable"
	keyCapabilities                      = "capabilities"
	keyCh                                = "ch"
	keyChange                            = "change"
	keyChangeNotifications               = "changeNotifications"
	keyChanges                           = "changes"
	keyCharacter                         = "character"
	keyChildren                          = "children"
	keyClientInfo                        = "clientInfo"
	keyCode                              = "code"
	keyCodeAction                        = "codeAction"
	keyCodeActionKind                    = "codeActionKind"
	keyCodeActionKinds                   = "codeActionKinds"
	keyCodeActionLiteralSupport          = "codeActionLiteralSupport"
	keyCodeActionProvider                = "codeActionProvider"
	keyCodeLens                          = "codeLens"
	keyCodeLensProvider                  = "codeLensProvider"
	keyColor                             = "color"
	keyColorProvider                     = "colorProvider"
	keyCommand                           = "command"
	keyCommands                          = "commands"
	keyCommitCharacters                  = "commitCharacters"
	keyCommitCharactersSupport           = "commitCharactersSupport"
	keyCompletion                        = "completion"
	keyCompletionItem                    = "completionItem"
	keyCompletionItemKind                = "completionItemKind"
	keyCompletionProvider                = "completionProvider"
	keyConfiguration                     = "configuration"
	keyContainerName                     = "containerName"
	keyContentChanges                    = "contentChanges"
	keyContentFormat                     = "contentFormat"
	keyContents                          = "contents"
	keyContext                           = "context"
	keyContextSupport                    = "contextSupport"
	keyData                              = "data"
	keyDeclaration                       = "declaration"
	keyDeclarationProvider               = "declarationProvider"
	keyDefinition                        = "definition"
	keyDefinitionProvider                = "definitionProvider"
	keyDeprecated                        = "deprecated"
	keyDeprecatedSupport                 = "deprecatedSupport"
	keyDetail                            = "detail"
	keyDiagnostics                       = "diagnostics"
	keyDidChangeConfiguration            = "didChangeConfiguration"
	keyDidChangeWatchedFiles             = "didChangeWatchedFiles"
	keyDidSave                           = "didSave"
	keyDocument                          = "document"
	keyDocumentation                     = "documentation"
	keyDocumentationFormat               = "documentationFormat"
	keyDocumentChanges                   = "documentChanges"
	keyDocumentFormattingProvider        = "documentFormattingProvider"
	keyDocumentHighlight                 = "documentHighlight"
	keyDocumentHighlightProvider         = "documentHighlightProvider"
	keyDocumentLink                      = "documentLink"
	keyDocumentLinkProvider              = "documentLinkProvider"
	keyDocumentOnTypeFormattingProvider  = "documentOnTypeFormattingProvider"
	keyDocumentRangeFormattingProvider   = "documentRangeFormattingProvider"
	keyDocumentSelector                  = "documentSelector"
	keyDocumentSymbol                    = "documentSymbol"
	keyDocumentSymbolProvider            = "documentSymbolProvider"
	keyDynamicRegistration               = "dynamicRegistration"
	keyEdit                              = "edit"
	keyEdits                             = "edits"
	keyEnd                               = "end"
	keyEndCharacter                      = "endCharacter"
	keyEndLine                           = "endLine"
	keyEvent                             = "event"
	keyExecuteCommand                    = "executeCommand"
	keyExecuteCommandProvider            = "executeCommandProvider"
	keyExperimental                      = "experimental"
	keyFailureHandling                   = "failureHandling"
	keyFilterText                        = "filterText"
	keyFirstTriggerCharacter             = "firstTriggerCharacter"
	keyFoldingRange                      = "foldingRange"
	keyFoldingRangeProvider              = "foldingRangeProvider"
	keyFormatting                        = "formatting"
	keyGlobPattern                       = "globPattern"
	keyGreen                             = "green"
	keyHierarchicalDocumentSymbolSupport = "hierarchicalDocumentSymbolSupport"
	keyHover                             = "hover"
	keyHoverProvider                     = "hoverProvider"
	keyID                                = "id"
	keyIgnoreIfExists                    = "ignoreIfExists"
	keyIgnoreIfNotExists                 = "ignoreIfNotExists"
	keyImplementation                    = "implementation"
	keyImplementationProvider            = "implementationProvider"
	keyIncludeDeclaration                = "includeDeclaration"
	keyIncludeText                       = "includeText"
	keyInitializationOptions             = "initializationOptions"
	keyInsertFinalNewline                = "insertFinalNewline"
	keyInsertSpaces                      = "insertSpaces"
	keyInsertText                        = "insertText"
	keyInsertTextFormat                  = "insertTextFormat"
	keyIsIncomplete                      = "isIncomplete"
	keyIsPreferred                       = "isPreferred"
	keyIsPreferredSupport                = "isPreferredSupport"
	keyIsRetrigger                       = "isRetrigger"
	keyItems                             = "items"
	keyKind                              = "kind"
	keyLabel                             = "label"
	keyLabelOffsetSupport                = "labelOffsetSupport"
	keyLanguage                          = "language"
	keyLanguageID                        = "languageId"
	keyLine                              = "line"
	keyLineCount                         = "lineCount"
	keyLineFoldingOnly                   = "lineFoldingOnly"
	keyLinkSupport                       = "linkSupport"
	keyLocation                          = "location"
	keyMessage                           = "message"
	keyMethod                            = "method"
	keyMoreTriggerCharacter              = "moreTriggerCharacter"
	keyName                              = "name"
	keyNewName                           = "newName"
	keyNewText                           = "newText"
	keyNewURI                            = "newUri"
	keyOldURI                            = "oldUri"
	keyOnly                              = "only"
	keyOnTypeFormatting                  = "onTypeFormatting"
	keyOpenClose                         = "openClose"
	keyOptions                           = "options"
	keyOriginSelectionRange              = "originSelectionRange"
	keyOverwrite                         = "overwrite"
	keyParameterInformation              = "parameterInformation"
	keyParent                            = "parent"
	keyPartialResultToken                = "partialResultToken"
	keyPattern                           = "pattern"
	keyPercentage                        = "percentage"
	keyPosition                          = "position"
	keyPositions                         = "positions"
	keyPrepareProvider                   = "prepareProvider"
	keyPrepareSupport                    = "prepareSupport"
	keyPreselect                         = "preselect"
	keyPreselectSupport                  = "preselectSupport"
	keyProcessID                         = "processId"
	keyPublishDiagnostics                = "publishDiagnostics"
	keyQuery                             = "query"
	keyRange                             = "range"
	keyRangeFormatting                   = "rangeFormatting"
	keyRangeLength                       = "rangeLength"
	keyRangeLimit                        = "rangeLimit"
	keyReason                            = "reason"
	keyRecursive                         = "recursive"
	keyRed                               = "red"
	keyReferences                        = "references"
	keyReferencesProvider                = "referencesProvider"
	keyRegisterOptions                   = "registerOptions"
	keyRegistrations                     = "registrations"
	keyRelatedInformation                = "relatedInformation"
	keyRemoved                           = "removed"
	keyRename                            = "rename"
	keyRenameProvider                    = "renameProvider"
	keyResolveProvider                   = "resolveProvider"
	keyResourceOperations                = "resourceOperations"
	keyRetriggerCharacters               = "retriggerCharacters"
	keyRetry                             = "retry"
	keyRootPath                          = "rootPath"
	keyRootURI                           = "rootUri"
	keySave                              = "save"
	keyScheme                            = "scheme"
	keyScopeURI                          = "scopeUri"
	keySection                           = "section"
	keySelectionRange                    = "selectionRange"
	keySelectionRangeProvider            = "selectionRangeProvider"
	keyServerInfo                        = "serverInfo"
	keySettings                          = "settings"
	keySeverity                          = "severity"
	keySignatureHelp                     = "signatureHelp"
	keySignatureHelpProvider             = "signatureHelpProvider"
	keySignatureInformation              = "signatureInformation"
	keySignatures                        = "signatures"
	keySnippetSupport                    = "snippetSupport"
	keySortText                          = "sortText"
	keySource                            = "source"
	keyStart                             = "start"
	keyStartCharacter                    = "startCharacter"
	keyStartLine                         = "startLine"
	keySupported                         = "supported"
	keySymbol                            = "symbol"
	keySymbolKind                        = "symbolKind"
	keySynchronization                   = "synchronization"
	keySyncKind                          = "syncKind"
	keyTabSize                           = "tabSize"
	keyTags                              = "tags"
	keyTagSupport                        = "tagSupport"
	keyTarget                            = "target"
	keyTargetRange                       = "targetRange"
	keyTargetSelectionRange              = "targetSelectionRange"
	keyTargetURI                         = "targetUri"
	keyText                              = "text"
	keyTextDocument                      = "textDocument"
	keyTextDocumentSync                  = "textDocumentSync"
	keyTextEdit                          = "textEdit"
	keyTitle                             = "title"
	keyToken                             = "token"
	keyTooltip                           = "tooltip"
	keyTooltipSupport                    = "tooltipSupport"
	keyTrace                             = "trace"
	keyTriggerCharacter                  = "triggerCharacter"
	keyTriggerCharacters                 = "triggerCharacters"
	keyTriggerKind                       = "triggerKind"
	keyTrimFinalNewlines                 = "trimFinalNewlines"
	keyTrimTrailingWhitespace            = "trimTrailingWhitespace"
	keyType                              = "type"
	keyTypeDefinition                    = "typeDefinition"
	keyTypeDefinitionProvider            = "typeDefinitionProvider"
	keyUnregisterations                  = "unregisterations"
	keyURI                               = "uri"
	keyValue                             = "value"
	keyValueSet                          = "valueSet"
	keyVersion                           = "version"
	keyVersionSupport                    = "versionSupport"
	keyWatchers                          = "watchers"
	keyWillSave                          = "willSave"
	keyWillSaveWaitUntil                 = "willSaveWaitUntil"
	keyWindow                            = "window"
	keyWorkDoneProgress                  = "workDoneProgress"
	keyWorkDoneToken                     = "workDoneToken"
	keyWorkspace                         = "workspace"
	keyWorkspaceEdit                     = "workspaceEdit"
	keyWorkspaceFolders                  = "workspaceFolders"
	keyWorkspaceSymbolProvider           = "workspaceSymbolProvider"
)
