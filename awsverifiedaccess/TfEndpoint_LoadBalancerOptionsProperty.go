package awsverifiedaccess


// Experimental.
type TfEndpoint_LoadBalancerOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#load_balancer_arn TfEndpoint#load_balancer_arn}.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#port TfEndpoint#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#port_range TfEndpoint#port_range}
	// Experimental.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#protocol TfEndpoint#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#subnet_ids TfEndpoint#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

