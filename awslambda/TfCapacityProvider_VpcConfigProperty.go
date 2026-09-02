package awslambda


// Experimental.
type TfCapacityProvider_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#security_group_ids TfCapacityProvider#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#subnet_ids TfCapacityProvider#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

