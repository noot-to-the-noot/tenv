/*
 *
 * Copyright 2024 tofuutils authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"fmt"

	"github.com/tofuutils/tenv/v3/config/cmdconst"
	"github.com/tofuutils/tenv/v3/versionmanager/tenvlib"
)

func main() {
	tenv, err := tenvlib.Make(tenvlib.AutoInstall, tenvlib.IgnoreEnv, tenvlib.DisableDisplay)
	if err != nil {
		fmt.Println("init failed :", err)

		return
	}

	err = tenv.DetectedCommandProxy(context.Background(), cmdconst.TofuName, "version")
	if err != nil {
		fmt.Println("proxy call failed :", err)
	}
}