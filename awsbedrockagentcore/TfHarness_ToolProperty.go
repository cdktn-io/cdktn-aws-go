package awsbedrockagentcore


// Experimental.
type TfHarness_ToolProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#type TfHarness#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#config TfHarness#config}
	// Experimental.
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#name TfHarness#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

