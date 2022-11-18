/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Bucket.
func (mg *Bucket) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccessKeyRef,
		Selector:     mg.Spec.ForProvider.AccessKeySelector,
		To: reference.To{
			List:    &KeyList{},
			Managed: &Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessKey")
	}
	mg.Spec.ForProvider.AccessKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessKeyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecretKeyRef,
		Selector:     mg.Spec.ForProvider.SecretKeySelector,
		To: reference.To{
			List:    &KeyList{},
			Managed: &Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretKey")
	}
	mg.Spec.ForProvider.SecretKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretKeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Object.
func (mg *Object) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccessKeyRef,
		Selector:     mg.Spec.ForProvider.AccessKeySelector,
		To: reference.To{
			List:    &KeyList{},
			Managed: &Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessKey")
	}
	mg.Spec.ForProvider.AccessKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessKeyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecretKeyRef,
		Selector:     mg.Spec.ForProvider.SecretKeySelector,
		To: reference.To{
			List:    &KeyList{},
			Managed: &Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretKey")
	}
	mg.Spec.ForProvider.SecretKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretKeyRef = rsp.ResolvedReference

	return nil
}
