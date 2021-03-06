/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	registryv1 "tkestack.io/tke/api/registry/v1"
)

// FakeChartGroups implements ChartGroupInterface
type FakeChartGroups struct {
	Fake *FakeRegistryV1
}

var chartgroupsResource = schema.GroupVersionResource{Group: "registry.tkestack.io", Version: "v1", Resource: "chartgroups"}

var chartgroupsKind = schema.GroupVersionKind{Group: "registry.tkestack.io", Version: "v1", Kind: "ChartGroup"}

// Get takes name of the chartGroup, and returns the corresponding chartGroup object, and an error if there is any.
func (c *FakeChartGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *registryv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(chartgroupsResource, name), &registryv1.ChartGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*registryv1.ChartGroup), err
}

// List takes label and field selectors, and returns the list of ChartGroups that match those selectors.
func (c *FakeChartGroups) List(ctx context.Context, opts v1.ListOptions) (result *registryv1.ChartGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(chartgroupsResource, chartgroupsKind, opts), &registryv1.ChartGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &registryv1.ChartGroupList{ListMeta: obj.(*registryv1.ChartGroupList).ListMeta}
	for _, item := range obj.(*registryv1.ChartGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chartGroups.
func (c *FakeChartGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(chartgroupsResource, opts))
}

// Create takes the representation of a chartGroup and creates it.  Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *FakeChartGroups) Create(ctx context.Context, chartGroup *registryv1.ChartGroup, opts v1.CreateOptions) (result *registryv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(chartgroupsResource, chartGroup), &registryv1.ChartGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*registryv1.ChartGroup), err
}

// Update takes the representation of a chartGroup and updates it. Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *FakeChartGroups) Update(ctx context.Context, chartGroup *registryv1.ChartGroup, opts v1.UpdateOptions) (result *registryv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(chartgroupsResource, chartGroup), &registryv1.ChartGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*registryv1.ChartGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChartGroups) UpdateStatus(ctx context.Context, chartGroup *registryv1.ChartGroup, opts v1.UpdateOptions) (*registryv1.ChartGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(chartgroupsResource, "status", chartGroup), &registryv1.ChartGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*registryv1.ChartGroup), err
}

// Delete takes name of the chartGroup and deletes it. Returns an error if one occurs.
func (c *FakeChartGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(chartgroupsResource, name), &registryv1.ChartGroup{})
	return err
}

// Patch applies the patch and returns the patched chartGroup.
func (c *FakeChartGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *registryv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(chartgroupsResource, name, pt, data, subresources...), &registryv1.ChartGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*registryv1.ChartGroup), err
}
