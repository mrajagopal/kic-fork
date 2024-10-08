// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/dos/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDosProtectedResources implements DosProtectedResourceInterface
type FakeDosProtectedResources struct {
	Fake *FakeAppprotectdosV1beta1
	ns   string
}

var dosprotectedresourcesResource = v1beta1.SchemeGroupVersion.WithResource("dosprotectedresources")

var dosprotectedresourcesKind = v1beta1.SchemeGroupVersion.WithKind("DosProtectedResource")

// Get takes name of the dosProtectedResource, and returns the corresponding dosProtectedResource object, and an error if there is any.
func (c *FakeDosProtectedResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DosProtectedResource, err error) {
	emptyResult := &v1beta1.DosProtectedResource{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(dosprotectedresourcesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.DosProtectedResource), err
}

// List takes label and field selectors, and returns the list of DosProtectedResources that match those selectors.
func (c *FakeDosProtectedResources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DosProtectedResourceList, err error) {
	emptyResult := &v1beta1.DosProtectedResourceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(dosprotectedresourcesResource, dosprotectedresourcesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DosProtectedResourceList{ListMeta: obj.(*v1beta1.DosProtectedResourceList).ListMeta}
	for _, item := range obj.(*v1beta1.DosProtectedResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dosProtectedResources.
func (c *FakeDosProtectedResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(dosprotectedresourcesResource, c.ns, opts))

}

// Create takes the representation of a dosProtectedResource and creates it.  Returns the server's representation of the dosProtectedResource, and an error, if there is any.
func (c *FakeDosProtectedResources) Create(ctx context.Context, dosProtectedResource *v1beta1.DosProtectedResource, opts v1.CreateOptions) (result *v1beta1.DosProtectedResource, err error) {
	emptyResult := &v1beta1.DosProtectedResource{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(dosprotectedresourcesResource, c.ns, dosProtectedResource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.DosProtectedResource), err
}

// Update takes the representation of a dosProtectedResource and updates it. Returns the server's representation of the dosProtectedResource, and an error, if there is any.
func (c *FakeDosProtectedResources) Update(ctx context.Context, dosProtectedResource *v1beta1.DosProtectedResource, opts v1.UpdateOptions) (result *v1beta1.DosProtectedResource, err error) {
	emptyResult := &v1beta1.DosProtectedResource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(dosprotectedresourcesResource, c.ns, dosProtectedResource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.DosProtectedResource), err
}

// Delete takes name of the dosProtectedResource and deletes it. Returns an error if one occurs.
func (c *FakeDosProtectedResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dosprotectedresourcesResource, c.ns, name, opts), &v1beta1.DosProtectedResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDosProtectedResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(dosprotectedresourcesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DosProtectedResourceList{})
	return err
}

// Patch applies the patch and returns the patched dosProtectedResource.
func (c *FakeDosProtectedResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DosProtectedResource, err error) {
	emptyResult := &v1beta1.DosProtectedResource{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(dosprotectedresourcesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.DosProtectedResource), err
}
