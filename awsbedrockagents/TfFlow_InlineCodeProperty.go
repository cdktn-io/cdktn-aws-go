package awsbedrockagents


// Experimental.
type TfFlow_InlineCodeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#code TfFlow#code}.
	// Experimental.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#language TfFlow#language}.
	// Experimental.
	Language *string `field:"required" json:"language" yaml:"language"`
}

