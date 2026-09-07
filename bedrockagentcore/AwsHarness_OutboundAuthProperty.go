package bedrockagentcore


// Experimental.
type AwsHarness_OutboundAuthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#aws_iam AwsHarness#aws_iam}.
	// Experimental.
	AwsIam interface{} `field:"optional" json:"awsIam" yaml:"awsIam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#none AwsHarness#none}.
	// Experimental.
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#oauth AwsHarness#oauth}
	// Experimental.
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

