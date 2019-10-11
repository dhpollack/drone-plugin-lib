// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/urfave/cli"
)

//---------------------------------------------------------------------
// Transport Flags
//---------------------------------------------------------------------

// Network contains options for connecting to the network.
type Network struct {
	// Context for making network requests.
	//
	// If `trace` logging is requested the context will use `httptrace` to
	// capture all network requests.
	Context context.Context

	// Client for making network requests.
	Client *http.Client
}

const networkSSLVerifyFlag = "transport.ssl-verify"

// networkFlags has the cli.Flags for the Transport.
func networkFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   networkSSLVerifyFlag,
			Usage:  "transport ssl verify",
			EnvVar: "PLUGIN_SSL_VERIFY",
		},
	}
}

// NetworkFromContext creates a Transport from the cli.Context.
func NetworkFromContext(ctx *cli.Context) Network {
	// Create the client
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	// Create the context
	context := context.Background()

	if ctx.String(logLevelFlag) == "trace" {
		context = traceHTTP(context)
	}

	return Network{
		Client:  client,
		Context: context,
	}
}