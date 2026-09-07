package bedrockagents


// Experimental.
type AwsPrompt_ToolSpecProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#name AwsPrompt#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#description AwsPrompt#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#input_schema AwsPrompt#input_schema}
	// Experimental.
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
}

