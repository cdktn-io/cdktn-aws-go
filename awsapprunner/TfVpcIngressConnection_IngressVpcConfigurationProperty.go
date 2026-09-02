package awsapprunner


// Experimental.
type TfVpcIngressConnection_IngressVpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_vpc_ingress_connection#vpc_endpoint_id TfVpcIngressConnection#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_vpc_ingress_connection#vpc_id TfVpcIngressConnection#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

