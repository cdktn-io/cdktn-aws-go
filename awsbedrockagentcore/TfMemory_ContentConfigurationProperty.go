package awsbedrockagentcore


// Experimental.
type TfMemory_ContentConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#type TfMemory#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#level TfMemory#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

