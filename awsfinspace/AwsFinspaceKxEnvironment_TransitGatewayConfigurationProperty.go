package awsfinspace


// Experimental.
type AwsFinspaceKxEnvironment_TransitGatewayConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#routable_cidr_space AwsFinspaceKxEnvironment#routable_cidr_space}.
	// Experimental.
	RoutableCidrSpace *string `field:"required" json:"routableCidrSpace" yaml:"routableCidrSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#transit_gateway_id AwsFinspaceKxEnvironment#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// attachment_network_acl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#attachment_network_acl_configuration AwsFinspaceKxEnvironment#attachment_network_acl_configuration}
	// Experimental.
	AttachmentNetworkAclConfiguration interface{} `field:"optional" json:"attachmentNetworkAclConfiguration" yaml:"attachmentNetworkAclConfiguration"`
}

