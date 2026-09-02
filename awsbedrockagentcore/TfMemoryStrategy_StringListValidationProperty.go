package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_StringListValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#allowed_values TfMemoryStrategy#allowed_values}.
	// Experimental.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#max_items TfMemoryStrategy#max_items}.
	// Experimental.
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

