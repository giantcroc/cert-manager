/*
Copyright 2020 The cert-manager Authors.

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

// Package v1alpha3 is the v1alpha3 version of the API.
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/jetstack/cert-manager/pkg/apis/experimental
// +k8s:openapi-gen=true
// +k8s:defaulter-gen=TypeMeta
// +groupName=experimental.cert-manager.io
// +groupGoName=Experimental
package v1alpha3
