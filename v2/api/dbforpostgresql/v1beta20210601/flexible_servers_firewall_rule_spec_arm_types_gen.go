// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServers_FirewallRule_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a firewall rule.
	Properties *FirewallRuleProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_FirewallRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (rule FlexibleServers_FirewallRule_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *FlexibleServers_FirewallRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rule *FlexibleServers_FirewallRule_Spec_ARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

type FirewallRuleProperties_ARM struct {
	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
