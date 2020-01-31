// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package test

// TestFetcherFactory is a type representing a function which builds
// a TestFetcher
type TestFetcherFactory func() TestFetcher

// TestFetcher is an interface used to get the test to run on the selected
// hosts.
type TestFetcher interface {
	ValidateFetchParameters([]byte) (interface{}, error)
	Fetch(interface{}) (string, []*TestStepDescriptor, error)
}

// TestFetcherBundle bundles the selected TestFetcher together with its acquire
// and release parameters based on the content of the job descriptor
type TestFetcherBundle struct {
	TestFetcher     TestFetcher
	FetchParameters interface{}
}
