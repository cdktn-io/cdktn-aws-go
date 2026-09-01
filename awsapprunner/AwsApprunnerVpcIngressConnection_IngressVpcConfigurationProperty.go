package awsapprunner


// Experimental.
type AwsApprunnerVpcIngressConnection_IngressVpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_vpc_ingress_connection#vpc_endpoint_id AwsApprunnerVpcIngressConnection#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_vpc_ingress_connection#vpc_id AwsApprunnerVpcIngressConnection#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

