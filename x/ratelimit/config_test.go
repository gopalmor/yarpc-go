// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ratelimit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/yarpc/internal/whitespace"

	yaml "gopkg.in/yaml.v2"
)

// func TestUnaryInboundMiddlewareConfig(t *testing.T) {
// 	given := whitespace.Expand(`
// 		rps: 10
// 		burstLimit: 10
// 	`)
// 	var unstructured config.AttributeMap
// 	err := yaml.Unmarshal([]byte(given), &unstructured)
// 	require.NoError(t, err)

// 	var config UnaryInboundMiddlewareConfig
// 	err = unstructured.Decode(&config)
// 	require.NoError(t, err)

// 	assert.Equal(t, 10, config.RPS)
// 	assert.Equal(t, 10, config.BurstLimit)
// 	assert.Equal(t, false, config.NoSlack)

// 	_, err = config.Build()
// 	require.NoError(t, err)
// }

// func TestUnaryInboundMiddlewareConfigWithoutSlack(t *testing.T) {
// 	given := whitespace.Expand(`
// 		rps: 10
// 		noSlack: true
// 	`)
// 	var unstructured config.AttributeMap
// 	err := yaml.Unmarshal([]byte(given), &unstructured)
// 	require.NoError(t, err)

// 	var config UnaryInboundMiddlewareConfig
// 	err = unstructured.Decode(&config)
// 	require.NoError(t, err)

// 	assert.Equal(t, 10, config.RPS)
// 	assert.Equal(t, 0, config.BurstLimit)
// 	assert.Equal(t, true, config.NoSlack)

// 	_, err = config.Build()
// 	require.NoError(t, err)
// }

// func TestUnaryInboundMiddlewareConfigContradiction(t *testing.T) {
// 	_, err := UnaryInboundMiddlewareConfig{
// 		NoSlack:    true,
// 		BurstLimit: 10,
// 	}.Build()
// 	require.Error(t, err)
// }

func TestConfiguration(t *testing.T) {
	ratelimitConfig := whitespace.Expand(`
	ratelimit:
	- rps: 10
		procedure: Hello::foo
	- rps: 500
    burstLimit: 10
		procedure: Hello::echo
	- rps: 10
		burstLimit: 5
		procedure: Hello::call
	- rps: 20
		procedure: Hello::bar
	- rps: 1000
		burstLimit: 10
		global: true
	- rps: 5
		burstLimit: 10
		default: true
	`)

	var cfg map[string][]InboundThrottleConfig
	err := yaml.Unmarshal([]byte(ratelimitConfig), &cfg)
	require.NoError(t, err)
	t.Fatal(cfg)
}
