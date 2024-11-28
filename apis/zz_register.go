/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/yandex-cloud/crossplane-provider-yc/apis/alb/v1alpha1"
	v1alpha1cdn "github.com/yandex-cloud/crossplane-provider-yc/apis/cdn/v1alpha1"
	v1alpha1compute "github.com/yandex-cloud/crossplane-provider-yc/apis/compute/v1alpha1"
	v1alpha1container "github.com/yandex-cloud/crossplane-provider-yc/apis/container/v1alpha1"
	v1alpha1datatransfer "github.com/yandex-cloud/crossplane-provider-yc/apis/datatransfer/v1alpha1"
	v1alpha1dns "github.com/yandex-cloud/crossplane-provider-yc/apis/dns/v1alpha1"
	v1alpha1iam "github.com/yandex-cloud/crossplane-provider-yc/apis/iam/v1alpha1"
	v1alpha1kms "github.com/yandex-cloud/crossplane-provider-yc/apis/kms/v1alpha1"
	v1alpha1kubernetes "github.com/yandex-cloud/crossplane-provider-yc/apis/kubernetes/v1alpha1"
	v1alpha1mdb "github.com/yandex-cloud/crossplane-provider-yc/apis/mdb/v1alpha1"
	v1alpha1message "github.com/yandex-cloud/crossplane-provider-yc/apis/message/v1alpha1"
	v1alpha1resourcemanager "github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1"
	v1alpha1storage "github.com/yandex-cloud/crossplane-provider-yc/apis/storage/v1alpha1"
	v1alpha1apis "github.com/yandex-cloud/crossplane-provider-yc/apis/v1alpha1"
	v1beta1 "github.com/yandex-cloud/crossplane-provider-yc/apis/v1beta1"
	v1alpha1vpc "github.com/yandex-cloud/crossplane-provider-yc/apis/vpc/v1alpha1"
	v1alpha1ydb "github.com/yandex-cloud/crossplane-provider-yc/apis/ydb/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1cdn.SchemeBuilder.AddToScheme,
		v1alpha1compute.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha1datatransfer.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1iam.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha1kubernetes.SchemeBuilder.AddToScheme,
		v1alpha1mdb.SchemeBuilder.AddToScheme,
		v1alpha1message.SchemeBuilder.AddToScheme,
		v1alpha1resourcemanager.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vpc.SchemeBuilder.AddToScheme,
		v1alpha1ydb.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}