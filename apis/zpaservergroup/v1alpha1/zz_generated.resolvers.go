/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/zscaler/crossplane-provider-zpa/apis/zpaappconnectorgroup/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AppConnectorGroups); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AppConnectorGroups[i3].ID),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.AppConnectorGroups[i3].IDRefs,
			Selector:      mg.Spec.ForProvider.AppConnectorGroups[i3].IDSelector,
			To: reference.To{
				List:    &v1alpha1.ConnectorGroupList{},
				Managed: &v1alpha1.ConnectorGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AppConnectorGroups[i3].ID")
		}
		mg.Spec.ForProvider.AppConnectorGroups[i3].ID = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.AppConnectorGroups[i3].IDRefs = mrsp.ResolvedReferences

	}

	return nil
}