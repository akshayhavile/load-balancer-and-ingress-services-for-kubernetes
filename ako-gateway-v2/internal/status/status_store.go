/*
 * Copyright 2021 VMware, Inc.
 * All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package status

import (
	"sync"

	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/internal/objects"

	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/utils"
)

//store ingress/route to vsname mapping
var aviStatusinstance *AviStatusLister
var avionce sync.Once

func SharedAviGraphLister() *AviStatusLister {
	avionce.Do(func() {
		AviStatusStore := objects.NewObjectMapStore()
		aviStatusinstance = &AviStatusLister{}
		aviStatusinstance.AviStatusStore = AviStatusStore
	})
	return aviStatusinstance
}

type AviStatusLister struct {
	AviStatusStore *objects.ObjectMapStore
}

func (a *AviStatusLister) Save(ingRouteName string, vsName interface{}) {
	utils.AviLog.Infof("Saving Model: %s", ingRouteName)
	a.AviStatusStore.AddOrUpdate(ingRouteName, vsName)
}

func (a *AviStatusLister) Get(ingRouteName string) (bool, interface{}) {
	ok, obj := a.AviStatusStore.Get(ingRouteName)
	return ok, obj
}

func (a *AviStatusLister) Delete(ingRouteName string) {
	a.AviStatusStore.Delete(ingRouteName)
}
