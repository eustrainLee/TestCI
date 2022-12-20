/*
Copyright 2021 The Everoute Authors.

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

package informer_test

import (
	"testing"

	"github.com/everoute/everoute/plugin/tower/pkg/informer"
	"github.com/everoute/everoute/plugin/tower/pkg/schema"
)

func TestGetTowerObjectKeyCrash(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("receive panic: %s, should not crash when receiving an empty object", err)
		}
	}()

	var emptyVM *schema.VM
	_, _ = informer.TowerObjectKey(emptyVM)
}
