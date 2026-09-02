package awslambda


// Experimental.
type TfFunction_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#security_group_ids TfFunction#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#subnet_ids TfFunction#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#ipv6_allowed_for_dual_stack TfFunction#ipv6_allowed_for_dual_stack}.
	// Experimental.
	Ipv6AllowedForDualStack interface{} `field:"optional" json:"ipv6AllowedForDualStack" yaml:"ipv6AllowedForDualStack"`
}

