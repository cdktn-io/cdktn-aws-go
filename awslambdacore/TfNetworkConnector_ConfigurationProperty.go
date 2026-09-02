package awslambdacore


// Experimental.
type TfNetworkConnector_ConfigurationProperty struct {
	// vpc_egress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdacore_network_connector#vpc_egress_configuration TfNetworkConnector#vpc_egress_configuration}
	// Experimental.
	VpcEgressConfiguration interface{} `field:"optional" json:"vpcEgressConfiguration" yaml:"vpcEgressConfiguration"`
}

