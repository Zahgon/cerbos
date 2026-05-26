// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package git

import (
	"context"
	"time"

	"github.com/go-git/go-git/v6/plumbing/client"

	"github.com/cerbos/cerbos/internal/storage"
)

const (
	confKey                 = storage.ConfKey + ".git"
	defaultOperationTimeout = 60 * time.Second
)

// Conf is required (if driver is set to 'git') configuration for Git storage driver.
// +desc=This section is required only if storage.driver is git.
type Conf struct {
	// SSH holds auth details for the SSH protocol.
	SSH *SSHAuth `yaml:"ssh,omitempty"`
	// HTTPS holds auth details for the HTTPS protocol.
	HTTPS *HTTPSAuth `yaml:"https,omitempty"`
	// OperationTimeout specifies the timeout for git operations.
	OperationTimeout *time.Duration `yaml:"operationTimeout,omitempty" conf:",example=60s"`
	// Protocol is the Git protocol to use. Valid values are https, ssh, and file.
	Protocol string `yaml:"protocol" conf:"required,example=file"`
	// URL is the URL to the Git repo.
	URL string `yaml:"url" conf:"required,example=file://${HOME}/tmp/cerbos/policies"`
	// Branch is the branch to checkout.
	Branch string `yaml:"branch" conf:",example=policies"`
	// SubDir is the path under the checked-out Git repo where the policies are stored.
	SubDir string `yaml:"subDir,omitempty" conf:",example=policies"`
	// CheckoutDir is the local path to checkout the Git repo to.
	CheckoutDir string `yaml:"checkoutDir" conf:",example=${HOME}/tmp/cerbos/work"`
	// [DEPRECATED] ScratchDir is the directory to use for holding temporary data.
	ScratchDir string `yaml:"scratchDir" conf:",ignore"`
	// UpdatePollInterval specifies the interval to poll the Git repository for changes. Set to 0 to disable.
	UpdatePollInterval time.Duration `yaml:"updatePollInterval" conf:",example=60s"`
}

// SSHAuth holds auth details for the SSH protocol.
type SSHAuth struct {
	// The git user. Defaults to git.
	User string `yaml:"user" conf:",example=git"`
	// The path to the SSH private key file.
	PrivateKeyFile string `yaml:"privateKeyFile" conf:",example=${HOME}/.ssh/id_rsa"`
	// The password to the SSH private key.
	Password string `yaml:"password" conf:",example=pw"` //nolint:gosec
}

func (sa *SSHAuth) Auth() (client.SSHAuth, error) {
	_ = "STUB: not implemented"
	return *new(client.SSHAuth), nil
}

// HTTPSAuth holds auth details for the HTTPS protocol.
type HTTPSAuth struct {
	// The username to use for authentication.
	Username string `yaml:"username" conf:",example=cerbos"`
	// The password (or token) to use for authentication.
	Password string `yaml:"password" conf:",example=${GITHUB_TOKEN}"` //nolint:gosec
}

func (ha *HTTPSAuth) Auth() (client.HTTPAuth, error) {
	_ = "STUB: not implemented"
	return *new(client.HTTPAuth), nil
}

func (conf *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) Validate() (errs error) { _ = "STUB: not implemented"; return nil }

func (conf *Conf) getSubDir() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) getBranch() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) getAuth() (client.Option, error) {
	_ = "STUB: not implemented"
	return *new(client.Option), nil
}

func (conf *Conf) getOpCtx(parent context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
