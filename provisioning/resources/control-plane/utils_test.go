/**
 * Copyright (c) 2016-2017 Snowplow Analytics Ltd.
 * All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache
 * License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at
 * http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied.
 *
 * See the Apache License Version 2.0 for the specific language
 * governing permissions and limitations there under.
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
