/*
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

package fake

import (
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrafficSplits implements TrafficSplitInterface
type FakeTrafficSplits struct {
	Fake *FakeSplitV1alpha3
	ns   string
}

var trafficsplitsResource = schema.GroupVersionResource{Group: "split.smi-spec.io", Version: "v1alpha3", Resource: "trafficsplits"}

var trafficsplitsKind = schema.GroupVersionKind{Group: "split.smi-spec.io", Version: "v1alpha3", Kind: "TrafficSplit"}

// Get takes name of the trafficSplit, and returns the corresponding trafficSplit object, and an error if there is any.
func (c *FakeTrafficSplits) Get(name string, options v1.GetOptions) (result *v1alpha3.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficsplitsResource, c.ns, name), &v1alpha3.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.TrafficSplit), err
}

// List takes label and field selectors, and returns the list of TrafficSplits that match those selectors.
func (c *FakeTrafficSplits) List(opts v1.ListOptions) (result *v1alpha3.TrafficSplitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficsplitsResource, trafficsplitsKind, c.ns, opts), &v1alpha3.TrafficSplitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.TrafficSplitList{ListMeta: obj.(*v1alpha3.TrafficSplitList).ListMeta}
	for _, item := range obj.(*v1alpha3.TrafficSplitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficSplits.
func (c *FakeTrafficSplits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficsplitsResource, c.ns, opts))

}

// Create takes the representation of a trafficSplit and creates it.  Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *FakeTrafficSplits) Create(trafficSplit *v1alpha3.TrafficSplit) (result *v1alpha3.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficsplitsResource, c.ns, trafficSplit), &v1alpha3.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.TrafficSplit), err
}

// Update takes the representation of a trafficSplit and updates it. Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *FakeTrafficSplits) Update(trafficSplit *v1alpha3.TrafficSplit) (result *v1alpha3.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficsplitsResource, c.ns, trafficSplit), &v1alpha3.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.TrafficSplit), err
}

// Delete takes name of the trafficSplit and deletes it. Returns an error if one occurs.
func (c *FakeTrafficSplits) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trafficsplitsResource, c.ns, name), &v1alpha3.TrafficSplit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficSplits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficsplitsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha3.TrafficSplitList{})
	return err
}

// Patch applies the patch and returns the patched trafficSplit.
func (c *FakeTrafficSplits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficsplitsResource, c.ns, name, pt, data, subresources...), &v1alpha3.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.TrafficSplit), err
}
