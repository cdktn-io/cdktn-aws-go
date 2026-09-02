package awsnetworkfirewall


// Experimental.
type TfVpcEndpointAssociation_SubnetMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_vpc_endpoint_association#subnet_id TfVpcEndpointAssociation#subnet_id}.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_vpc_endpoint_association#ip_address_type TfVpcEndpointAssociation#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

