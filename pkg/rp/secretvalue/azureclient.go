/*
Copyright 2023 The Radius Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secretvalue

import (
	"context"
	"fmt"

	"github.com/go-openapi/jsonpointer"
	"github.com/project-radius/radius/pkg/azure/armauth"
	"github.com/project-radius/radius/pkg/azure/clientv2"
	"github.com/project-radius/radius/pkg/resourcemodel"
	"github.com/project-radius/radius/pkg/ucp/resources"
	"github.com/project-radius/radius/pkg/ucp/store"
)

func NewSecretValueClient(arm *armauth.ArmConfig) SecretValueClient {
	return &client{ARM: arm}
}

var _ SecretValueClient = (*client)(nil)

type client struct {
	ARM *armauth.ArmConfig
}

func (c *client) FetchSecret(ctx context.Context, identity resourcemodel.ResourceIdentity, action string, valueSelector string) (any, error) {
	arm := &resourcemodel.ARMIdentity{}
	if err := store.DecodeMap(identity.Data, arm); err != nil {
		return nil, fmt.Errorf("unsupported resource type: %+v. Currently only ARM resources are supported", identity)
	}

	parsed, err := resources.ParseResource(arm.ID)
	if err != nil {
		return nil, err
	}

	client, err := clientv2.NewCustomActionClient(parsed.FindScope(resources.SubscriptionsSegment), &c.ARM.ClientOptions)
	if err != nil {
		return nil, err
	}

	response, err := client.InvokeCustomAction(ctx, arm.ID, arm.APIVersion, action)
	if err != nil {
		return nil, err
	}

	pointer, err := jsonpointer.New(valueSelector)
	if err != nil {
		return nil, err
	}

	secretValue, _, err := pointer.Get(response.Body)
	if err != nil {
		return nil, err
	}

	return secretValue, err
}
