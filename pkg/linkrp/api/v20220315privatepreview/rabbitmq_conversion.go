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

package v20220315privatepreview

import (
	"fmt"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
	"github.com/project-radius/radius/pkg/to"
)

// ConvertTo converts from the versioned RabbitMQMessageQueue resource to version-agnostic datamodel.
func (src *RabbitMQMessageQueueResource) ConvertTo() (v1.DataModelInterface, error) {
	converted := &datamodel.RabbitMQMessageQueue{
		BaseResource: v1.BaseResource{
			TrackedResource: v1.TrackedResource{
				ID:       to.String(src.ID),
				Name:     to.String(src.Name),
				Type:     to.String(src.Type),
				Location: to.String(src.Location),
				Tags:     to.StringMap(src.Tags),
			},
			InternalMetadata: v1.InternalMetadata{
				UpdatedAPIVersion:      Version,
				AsyncProvisioningState: toProvisioningStateDataModel(src.Properties.GetRabbitMQMessageQueueProperties().ProvisioningState),
			},
		},
		Properties: datamodel.RabbitMQMessageQueueProperties{
			BasicResourceProperties: rpv1.BasicResourceProperties{
				Environment: to.String(src.Properties.GetRabbitMQMessageQueueProperties().Environment),
				Application: to.String(src.Properties.GetRabbitMQMessageQueueProperties().Application),
			},
		},
	}
	switch v := src.Properties.(type) {
	case *ValuesRabbitMQMessageQueueProperties:
		if v.Queue == nil {
			return nil, v1.NewClientErrInvalidRequest("queue is a required property for mode 'values'")
		}
		converted.Properties.Queue = to.String(v.Queue)
		converted.Properties.Mode = datamodel.LinkModeValues
	case *RecipeRabbitMQMessageQueueProperties:
		if v.Recipe == nil {
			return nil, v1.NewClientErrInvalidRequest("recipe is a required property for mode 'recipe'")
		}
		converted.Properties.Recipe = toRecipeDataModel(v.Recipe)
		converted.Properties.Queue = to.String(v.Queue)
		converted.Properties.Mode = datamodel.LinkModeRecipe
	default:
		return nil, v1.NewClientErrInvalidRequest(fmt.Sprintf("Unsupported mode %s", *src.Properties.GetRabbitMQMessageQueueProperties().Mode))
	}
	if src.Properties.GetRabbitMQMessageQueueProperties().Secrets != nil {
		converted.Properties.Secrets = datamodel.RabbitMQSecrets{
			ConnectionString: to.String(src.Properties.GetRabbitMQMessageQueueProperties().Secrets.ConnectionString),
		}
	}
	return converted, nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned RabbitMQMessageQueue resource.
func (dst *RabbitMQMessageQueueResource) ConvertFrom(src v1.DataModelInterface) error {
	rabbitmq, ok := src.(*datamodel.RabbitMQMessageQueue)
	if !ok {
		return v1.ErrInvalidModelConversion
	}

	dst.ID = to.Ptr(rabbitmq.ID)
	dst.Name = to.Ptr(rabbitmq.Name)
	dst.Type = to.Ptr(rabbitmq.Type)
	dst.SystemData = fromSystemDataModel(rabbitmq.SystemData)
	dst.Location = to.Ptr(rabbitmq.Location)
	dst.Tags = *to.StringMapPtr(rabbitmq.Tags)
	switch rabbitmq.Properties.Mode {
	case datamodel.LinkModeValues:
		mode := "values"
		dst.Properties = &ValuesRabbitMQMessageQueueProperties{
			Status: &ResourceStatus{
				OutputResources: rpv1.BuildExternalOutputResources(rabbitmq.Properties.Status.OutputResources),
			},
			ProvisioningState: fromProvisioningStateDataModel(rabbitmq.InternalMetadata.AsyncProvisioningState),
			Environment:       to.Ptr(rabbitmq.Properties.Environment),
			Application:       to.Ptr(rabbitmq.Properties.Application),
			Mode:              &mode,
			Queue:             to.Ptr(rabbitmq.Properties.Queue),
		}
	case datamodel.LinkModeRecipe:
		mode := "recipe"
		var recipe *Recipe
		recipe = fromRecipeDataModel(rabbitmq.Properties.Recipe)
		dst.Properties = &RecipeRabbitMQMessageQueueProperties{
			Status: &ResourceStatus{
				OutputResources: rpv1.BuildExternalOutputResources(rabbitmq.Properties.Status.OutputResources),
			},
			ProvisioningState: fromProvisioningStateDataModel(rabbitmq.InternalMetadata.AsyncProvisioningState),
			Environment:       to.Ptr(rabbitmq.Properties.Environment),
			Application:       to.Ptr(rabbitmq.Properties.Application),
			Mode:              &mode,
			Queue:             to.Ptr(rabbitmq.Properties.Queue),
			Recipe:            recipe,
		}
	default:
		return fmt.Errorf("unsupported mode %s", rabbitmq.Properties.Mode)
	}
	return nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned RabbitmqSecrets instance.
func (dst *RabbitMQSecrets) ConvertFrom(src v1.DataModelInterface) error {
	rabbitMQSecrets, ok := src.(*datamodel.RabbitMQSecrets)
	if !ok {
		return v1.ErrInvalidModelConversion
	}

	dst.ConnectionString = to.Ptr(rabbitMQSecrets.ConnectionString)
	return nil
}

// ConvertTo converts from the versioned RabbitMQSecrets instance to version-agnostic datamodel.
func (src *RabbitMQSecrets) ConvertTo() (v1.DataModelInterface, error) {
	converted := &datamodel.RabbitMQSecrets{
		ConnectionString: to.String(src.ConnectionString),
	}
	return converted, nil
}
