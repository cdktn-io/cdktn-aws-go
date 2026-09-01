package awsbedrockagents


// Experimental.
type AwsBedrockagentPrompt_MessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#role AwsBedrockagentPrompt#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#content AwsBedrockagentPrompt#content}
	// Experimental.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

