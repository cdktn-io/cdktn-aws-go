package bedrockagents


// Experimental.
type AwsPrompt_MessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#role AwsPrompt#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#content AwsPrompt#content}
	// Experimental.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

