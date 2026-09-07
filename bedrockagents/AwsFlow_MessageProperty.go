package bedrockagents


// Experimental.
type AwsFlow_MessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#role AwsFlow#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#content AwsFlow#content}
	// Experimental.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

