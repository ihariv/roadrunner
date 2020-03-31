// MIT License
//
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	rr "github.com/ihariv/roadrunner/cmd/rr/cmd"

	// services (plugins)
	"github.com/ihariv/roadrunner/service/env"
	"github.com/ihariv/roadrunner/service/gzip"
	"github.com/ihariv/roadrunner/service/headers"
	"github.com/ihariv/roadrunner/service/health"
	"github.com/ihariv/roadrunner/service/http"
	"github.com/ihariv/roadrunner/service/limit"
	"github.com/ihariv/roadrunner/service/metrics"
	"github.com/ihariv/roadrunner/service/reload"
	"github.com/ihariv/roadrunner/service/rpc"
	"github.com/ihariv/roadrunner/service/static"

	// additional commands and debug handlers
	_ "github.com/ihariv/roadrunner/cmd/rr/http"
	_ "github.com/ihariv/roadrunner/cmd/rr/limit"
)

func main() {
	rr.Container.Register(env.ID, &env.Service{})
	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(http.ID, &http.Service{})
	rr.Container.Register(metrics.ID, &metrics.Service{})
	rr.Container.Register(headers.ID, &headers.Service{})
	rr.Container.Register(static.ID, &static.Service{})
	rr.Container.Register(limit.ID, &limit.Service{})
	rr.Container.Register(health.ID, &health.Service{})
	rr.Container.Register(gzip.ID, &gzip.Service{})
	rr.Container.Register(reload.ID, &reload.Service{})

	// you can register additional commands using cmd.CLI
	rr.Execute()
}
