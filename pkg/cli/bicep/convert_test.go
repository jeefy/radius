// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package bicep

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ConvertMapStringInterface(t *testing.T) {
	in := make(map[string]map[string]interface{})
	in["throughput"] = map[string]interface{}{"value": 400}
	expected := map[string]interface{}{"throughput": 400}
	result := ConvertToMapStringInterface(in)
	require.Equal(t, result, expected)
}
