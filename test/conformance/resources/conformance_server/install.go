/*
 * Copyright 2021 The Knative Authors
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
 */

package conformance_server

import (
	"context"
	"embed"

	"knative.dev/reconciler-test/pkg/environment"
	"knative.dev/reconciler-test/pkg/feature"
	"knative.dev/reconciler-test/pkg/manifest"
)

//go:embed *.yaml
var yamls embed.FS

func StartPod(name string, port int, tls bool) feature.StepFn {
	cfg := map[string]interface{}{
		"name": name,
		"port": port,
		"tls":  tls,
	}

	return func(ctx context.Context, t feature.T) {
		if err := registerImage(ctx); err != nil {
			t.Fatal(err)
		}
		if _, err := manifest.InstallYamlFS(ctx, yamls, cfg); err != nil {
			t.Fatal(err)
		}
	}
}

func registerImage(ctx context.Context) error {
	im := manifest.ImagesFromFS(ctx, yamls)
	reg := environment.RegisterPackage(im...)
	_, err := reg(ctx, environment.FromContext(ctx))
	return err
}
