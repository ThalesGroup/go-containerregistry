// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package publish

import (
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
)

// WithTransport is a functional option for overriding the default transport
// on a default publisher.
func WithTransport(t http.RoundTripper) Option {
	return func(i *defaultOpener) error {
		i.setTransport(t)
		return nil
	}
}

// WithAuth is a functional option for overriding the default authenticator
// on a default publisher.
func WithAuth(auth authn.Authenticator) Option {
	return func(i *defaultOpener) error {
		i.setAuth(auth)
		return nil
	}
}

// WithAuthFromKeychain is a functional option for overriding the default
// authenticator on a default publisher using an authn.Keychain
func WithAuthFromKeychain(keys authn.Keychain) Option {
	return func(i *defaultOpener) error {
		auth, err := keys.Resolve(i.base.Registry)
		if err != nil {
			return err
		}
		i.setAuth(auth)
		return nil
	}
}

// Set client on image using provided transport, and the default authenticator
func (i *defaultOpener) setTransport(t http.RoundTripper) {
	i.t = t
}

// Set client on image using provided authenticator, and the default transport
func (i *defaultOpener) setAuth(auth authn.Authenticator) {
	i.auth = auth
}