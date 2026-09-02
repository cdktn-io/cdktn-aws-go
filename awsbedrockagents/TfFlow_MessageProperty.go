package awsbedrockagents


// Experimental.
type TfFlow_MessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#role TfFlow#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#content TfFlow#content}
	// Experimental.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

