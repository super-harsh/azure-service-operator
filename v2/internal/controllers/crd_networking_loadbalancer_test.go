/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_LoadBalancer_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Public IP Address
	sku := network.PublicIPAddressSkuNameStandard
	publicIPAddress := &network.PublicIPAddress{
		TypeMeta: metav1.TypeMeta{
			Kind: reflect.TypeOf(network.PublicIPAddress{}).Name(),
		},
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	// LoadBalancer
	loadBalancerSku := network.LoadBalancerSkuNameStandard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.InboundNatPoolPropertiesFormatProtocolTcp

	// TODO: This is still really awkward
	frontendIPConfigurationARMID, err := genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"loadBalancers",
		lbName,
		"frontendIPConfigurations",
		lbFrontendName)
	if err != nil {
		panic(err)
	}

	loadBalancer := &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.LoadBalancers_Spec_Properties_FrontendIPConfigurations{
				{
					Name: lbFrontendName,
					PublicIPAddress: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			// TODO: The below stuff isn't really necessary for LB CRUD but is required for VMSS...
			InboundNatPools: []network.LoadBalancers_Spec_Properties_InboundNatPools{
				{
					Name: "MyFancyNatPool",
					FrontendIPConfiguration: &network.SubResource{
						Reference: genruntime.ResourceReference{
							ARMID: frontendIPConfigurationARMID,
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.IntPtr(50_000),
					FrontendPortRangeEnd:   to.IntPtr(51_000),
					BackendPort:            to.IntPtr(22),
				},
			},
		},
	}

	tc.CreateResourceAndWait(loadBalancer)

	// It should be created in Kubernetes
	g.Expect(loadBalancer.Status.Id).ToNot(BeNil())
	armId := *loadBalancer.Status.Id

	tc.DeleteResourceAndWait(loadBalancer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(ctx, armId, string(network.LoadBalancersSpecAPIVersion20201101))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
