// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"encoding/json"
	"io"

	"github.com/cerbos/cerbos-sdk-go/cerbos/hub"
	storev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/store/v1"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

const (
	dirMode                = 0o700
	downloadBatchSize      = 10
	maxFileSize            = 5 * 1024 * 1024
	modifyFilesBatchSize   = 25
	replaceFilesZipMaxSize = 15728640
)

const storeCmdHelp = `Interact with Cerbos Hub managed stores.

Requires an existing managed store and the API credentials to access it.
The store ID and credentials can be provided using either command-line flags or environment variables.
`

type Conn struct {
	APIEndpoint   string `name:"api-endpoint" default:"https://api.cerbos.cloud" env:"CERBOS_HUB_API_ENDPOINT"`
	TLSCACert     string `name:"tls-ca-cert" hidden:"" help:"Path to the CA certificate for verifying server identity" type:"existingfile" env:"CERBOS_HUB_TLS_CA_CERT"`
	TLSClientCert string `name:"tls-client-cert" hidden:"" help:"Path to the TLS client certificate" type:"existingfile" env:"CERBOS_HUB_TLS_CLIENT_CERT" and:"tls-client-key"`
	TLSClientKey  string `name:"tls-client-key" hidden:"" help:"Path to the TLS client key" type:"existingfile" env:"CERBOS_HUB_TLS_CLIENT_KEY" and:"tls-client-cert"`
	StoreID       string `name:"store-id" help:"ID of the store to operate on" env:"CERBOS_HUB_STORE_ID" required:""`
	ClientID      string `name:"client-id" help:"Client ID of the access credential" env:"CERBOS_HUB_CLIENT_ID"`
	ClientSecret  string `name:"client-secret" help:"Client secret of the access credential" env:"CERBOS_HUB_CLIENT_SECRET"` //nolint:gosec
	TLSInsecure   bool   `name:"tls-insecure" hidden:"" help:"Skip validating server certificate" env:"CERBOS_HUB_TLS_INSECURE"`
}

func (c Conn) storeClient() (*hub.StoreClient, error) { _ = "STUB: not implemented"; return nil, nil }

type Cmd struct { //betteralign:ignore
	UploadGit    UploadGitCmd `cmd:"" name:"upload-git" help:"Upload files from a local git repository to the store"`
	Conn         `embed:""`
	ListFiles    ListFilesCmd    `cmd:"" name:"list-files" help:"List store files"`
	Download     DownloadCmd     `cmd:"" name:"download" help:"Download the entire store"`
	GetFiles     GetFilesCmd     `cmd:"" name:"get-files" help:"Download files from the store"`
	ReplaceFiles ReplaceFilesCmd `cmd:"" name:"replace-files" help:"Overwrite the store with the given set of files"`
	AddFiles     AddFilesCmd     `cmd:"" name:"add-files" help:"Add files to the store"`
	DeleteFiles  DeleteFilesCmd  `cmd:"" name:"delete-files" help:"Delete files from the store"`
}

func (*Cmd) Help() string { _ = "STUB: not implemented"; return "" }

type Output struct {
	Format string `name:"output" short:"o" default:"text" help:"Output format." enum:"text,json,prettyjson"`
}

func (o Output) toCommandError(w io.Writer, err error) error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

//nolint:mnd

func newNoFilesDownloadedError() error {
	_ = "STUB: not implemented"
	//nolint:mnd
	return nil
}

func (o Output) printNewVersion(w io.Writer, version int64) { _ = "STUB: not implemented"; return }

func (o Output) format(w io.Writer, value any) { _ = "STUB: not implemented"; return }

type ChangeDetails struct {
	Message  string          `help:"Commit message for this change"`
	Origin   json.RawMessage `help:"Metadata of the origin for this change as JSON string" placeholder:"{\"internal\":{\"source\":\"CI workflow\",\"metadata\":{\"id\":\"1\"}}}"`
	Uploader json.RawMessage `help:"Metadata of the uploader for this change as JSON string" placeholder:"{\"name\":\"cerbos-sdk-go\",\"metadata\":{\"version\":\"v0.1\"}}"`
}

func (cd ChangeDetails) ChangeDetails(gitChangeDetails *changeDetails) (*hub.ChangeDetails, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func changeDetailsFromHash(r *git.Repository, hash plumbing.Hash) (*changeDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type changeDetails struct {
	uploader *storev1.ChangeDetails_Uploader
	origin   *storev1.ChangeDetails_Git
	message  string
}

type commandError struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorDetails []any  `json:"errorDetails,omitempty"`
	exitCode     int
}

func (ce commandError) Error() string { _ = "STUB: not implemented"; return "" }

func (ce commandError) String() string { _ = "STUB: not implemented"; return "" }

func (ce commandError) ExitCode() int { _ = "STUB: not implemented"; return 0 }

type violation struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (v violation) String() string { _ = "STUB: not implemented"; return "" }

type validationErr struct {
	File    string `json:"file,omitempty"`
	Cause   string `json:"cause,omitempty"`
	Details string `json:"details,omitempty"`
}

func (ve validationErr) String() string { _ = "STUB: not implemented"; return "" }
