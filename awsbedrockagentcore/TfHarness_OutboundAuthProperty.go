package awsbedrockagentcore


// Experimental.
type TfHarness_OutboundAuthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#aws_iam TfHarness#aws_iam}.
	// Experimental.
	AwsIam interface{} `field:"optional" json:"awsIam" yaml:"awsIam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#none TfHarness#none}.
	// Experimental.
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#oauth TfHarness#oauth}
	// Experimental.
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

