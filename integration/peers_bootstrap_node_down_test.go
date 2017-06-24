// +build integration_disabled

// Copyright (c) 2016 Uber Technologies, Inc.
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

package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/m3db/m3db/integration/generate"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3db/storage/namespace"
	xlog "github.com/m3db/m3x/log"
)

func TestPeersBootstrapNodeDown(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	// Test setups
	log := xlog.SimpleLogger
	retentionOpts := retention.NewOptions().
		SetRetentionPeriod(6 * time.Hour).
		SetBlockSize(2 * time.Hour).
		SetBufferPast(10 * time.Minute).
		SetBufferFuture(2 * time.Minute)

	namesp := namespace.NewMetadata(testNamespaces[0],
		namespace.NewOptions().SetRetentionOptions(retentionOpts))
	opts := newTestOptions().
		SetNamespaces([]namespace.Metadata{namesp})

	setupOpts := []bootstrappableTestSetupOptions{
		{disablePeersBootstrapper: true},
		{disablePeersBootstrapper: true},
		{disablePeersBootstrapper: false},
	}
	setups, closeFn := newDefaultBootstrappableTestSetups(t, opts, setupOpts)
	defer closeFn()

	// Write test data for first node
	now := setups[0].getNowFn()
	blockSize := retentionOpts.BlockSize()
	seriesMaps := generate.BlocksByStart([]generate.BlockConfig{
		{[]string{"foo", "bar"}, 180, now.Add(-blockSize)},
		{[]string{"foo", "baz"}, 90, now},
	})
	err := writeTestDataToDisk(namesp, setups[0], seriesMaps)
	require.NoError(t, err)

	// Start the first server with filesystem bootstrapper
	require.NoError(t, setups[0].startServer())

	// Leave second node down, start the last server with peers and filesystem bootstrappers
	require.NoError(t, setups[2].startServer())
	log.Debug("first and third servers are now up")

	// Stop the servers
	defer func() {
		testSetups{setups[0], setups[2]}.parallel(func(s *testSetup) {
			require.NoError(t, s.stopServer())
		})
		log.Debug("servers are now down")
	}()

	// Verify in-memory data match what we expect
	expect := testSetups{setups[0], setups[2]}
	for _, setup := range expect {
		verifySeriesMaps(t, setup, namesp.ID(), seriesMaps)
	}
}
