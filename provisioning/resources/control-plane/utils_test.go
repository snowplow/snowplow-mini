/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.1
 * located at https://docs.snowplow.io/limited-use-license-1.1
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsJSON(t *testing.T) {
	assert := assert.New(t)

	validJSON := `{"test_key1":"test_value1", "test_key2":"test_value2"}`
	invalidJSON := `{"test_key1":"test_value1" "test_key2":"test_value2"}`

	assert.True(isJSON(validJSON))
	assert.False(isJSON(invalidJSON))
}

func TestIsURLReachable(t *testing.T) {
	assert := assert.New(t)

	reachableURL := "http://snowplowanalytics.com"
	unreachableURL := "http://unreachableurl.xyz"

	assert.True(isURLReachable(reachableURL))
	assert.False(isURLReachable(unreachableURL))
}

func TestIsValidUUID(t *testing.T) {
	assert := assert.New(t)

	validUUID := "29908bee-ff7a-4066-b724-bebe9da0e79a"
	invalidUUID := "7db4920d-07de-4110-a975-9f4480c8d6b"

	assert.True(isValidUUID(validUUID))
	assert.False(isValidUUID(invalidUUID))
}
