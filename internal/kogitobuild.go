// Copyright 2021 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"github.com/kiegroup/kogito-operator/api"
	"github.com/kiegroup/kogito-operator/core/client/kubernetes"
	"github.com/kiegroup/kogito-operator/core/kogitobuild"
	"github.com/kiegroup/kogito-operator/core/manager"
	rhpamv1 "github.com/kiegroup/rhpam-kogito-operator/api/v1"
	"k8s.io/apimachinery/pkg/types"
)

type kogitoBuildHandler struct {
	kogitobuild.BuildContext
}

// NewKogitoBuildHandler ...
func NewKogitoBuildHandler(buildContext kogitobuild.BuildContext) manager.KogitoBuildHandler {
	return &kogitoBuildHandler{
		BuildContext: buildContext,
	}
}

func (k *kogitoBuildHandler) FetchKogitoBuildInstance(key types.NamespacedName) (api.KogitoBuildInterface, error) {
	instance := &rhpamv1.KogitoBuild{}
	if exists, err := kubernetes.ResourceC(k.Client).FetchWithKey(key, instance); err != nil {
		return nil, err
	} else if !exists {
		return nil, nil
	}
	return instance, nil
}
