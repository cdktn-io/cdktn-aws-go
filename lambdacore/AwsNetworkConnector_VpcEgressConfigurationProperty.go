package lambdacore


// Experimental.
type AwsNetworkConnector_VpcEgressConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdacore_network_connector#associated_compute_resource_types AwsNetworkConnector#associated_compute_resource_types}.
	// Experimental.
	AssociatedComputeResourceTypes *[]*string `field:"required" json:"associatedComputeResourceTypes" yaml:"associatedComputeResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdacore_network_connector#security_group_ids AwsNetworkConnector#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdacore_network_connector#subnet_ids AwsNetworkConnector#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdacore_network_connector#network_protocol AwsNetworkConnector#network_protocol}.
	// Experimental.
	NetworkProtocol *string `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
}

