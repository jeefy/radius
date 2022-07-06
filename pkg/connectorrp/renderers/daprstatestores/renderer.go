// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package daprstatestores

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel"
	"github.com/project-radius/radius/pkg/connectorrp/renderers"
	"github.com/project-radius/radius/pkg/radrp/outputresource"
	"github.com/project-radius/radius/pkg/resourcekinds"
	"github.com/project-radius/radius/pkg/rp"
)

type StateStoreFunc = func(conv.DataModelInterface) ([]outputresource.OutputResource, error)

var SupportedStateStoreKindValues = map[string]StateStoreFunc{
	resourcekinds.DaprStateStoreAzureTableStorage: GetDaprStateStoreAzureStorage,
	resourcekinds.DaprGeneric:                     GetDaprStateStoreGeneric,
}

var _ renderers.Renderer = (*Renderer)(nil)

type Renderer struct {
	StateStores map[string]StateStoreFunc
}

func (r *Renderer) Render(ctx context.Context, dm conv.DataModelInterface) (renderers.RendererOutput, error) {
	resource, ok := dm.(datamodel.DaprStateStore)
	if !ok {
		return renderers.RendererOutput{}, conv.ErrInvalidModelConversion
	}

	properties := resource.Properties

	if r.StateStores == nil {
		return renderers.RendererOutput{}, errors.New("must support either kubernetes or ARM")
	}

	stateStoreFunc := r.StateStores[string(properties.Kind)]
	if stateStoreFunc == nil {
		return renderers.RendererOutput{}, fmt.Errorf("%s is not supported. Supported kind values: %s", properties.Kind, getAlphabeticallySortedKeys(r.StateStores))
	}

	resoures, err := stateStoreFunc(resource)
	if err != nil {
		return renderers.RendererOutput{}, err
	}

	values := map[string]renderers.ComputedValueReference{
		"stateStoreName": {
			Value: resource.Name,
		},
	}
	secrets := map[string]rp.SecretValueReference{}

	return renderers.RendererOutput{
		Resources:      resoures,
		ComputedValues: values,
		SecretValues:   secrets,
	}, nil
}

func getAlphabeticallySortedKeys(store map[string]StateStoreFunc) []string {
	keys := make([]string, len(store))

	i := 0
	for k := range store {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}
