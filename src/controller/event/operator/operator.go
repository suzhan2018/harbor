// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"context"

	"github.com/goharbor/harbor/src/common/security"
	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/pkg/user"
)

// FromContext return the event operator from context
func FromContext(ctx context.Context) string {
	sc, ok := security.FromContext(ctx)
	if !ok {
		return ""
	}

	if sc.IsSolutionUser() {
		user, err := user.Mgr.Get(ctx, 1)
		if err == nil {
			return user.Username
		}
		log.G(ctx).Errorf("failed to get operator for security %s, error: %v", sc.Name(), err)
	}

	return sc.GetUsername()
}
