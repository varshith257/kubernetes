/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"

	v1alpha2 "k8s.io/api/resource/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	resourcev1alpha2 "k8s.io/client-go/applyconfigurations/resource/v1alpha2"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ResourceClassParametersGetter has a method to return a ResourceClassParametersInterface.
// A group's client should implement this interface.
type ResourceClassParametersGetter interface {
	ResourceClassParameters(namespace string) ResourceClassParametersInterface
}

// ResourceClassParametersInterface has methods to work with ResourceClassParameters resources.
type ResourceClassParametersInterface interface {
	Create(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.CreateOptions) (*v1alpha2.ResourceClassParameters, error)
	Update(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.UpdateOptions) (*v1alpha2.ResourceClassParameters, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ResourceClassParameters, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ResourceClassParametersList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClassParameters, err error)
	Apply(ctx context.Context, resourceClassParameters *resourcev1alpha2.ResourceClassParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClassParameters, err error)
	ResourceClassParametersExpansion
}

// resourceClassParameters implements ResourceClassParametersInterface
type resourceClassParameters struct {
	*gentype.ClientWithListAndApply[*v1alpha2.ResourceClassParameters, *v1alpha2.ResourceClassParametersList, *resourcev1alpha2.ResourceClassParametersApplyConfiguration]
}

// newResourceClassParameters returns a ResourceClassParameters
func newResourceClassParameters(c *ResourceV1alpha2Client, namespace string) *resourceClassParameters {
	return &resourceClassParameters{
		gentype.NewClientWithListAndApply[*v1alpha2.ResourceClassParameters, *v1alpha2.ResourceClassParametersList, *resourcev1alpha2.ResourceClassParametersApplyConfiguration](
			"resourceclassparameters",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha2.ResourceClassParameters { return &v1alpha2.ResourceClassParameters{} },
			func() *v1alpha2.ResourceClassParametersList { return &v1alpha2.ResourceClassParametersList{} }),
	}
}
