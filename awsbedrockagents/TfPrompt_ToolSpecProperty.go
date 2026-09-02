package awsbedrockagents


// Experimental.
type TfPrompt_ToolSpecProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#name TfPrompt#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#description TfPrompt#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#input_schema TfPrompt#input_schema}
	// Experimental.
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
}

