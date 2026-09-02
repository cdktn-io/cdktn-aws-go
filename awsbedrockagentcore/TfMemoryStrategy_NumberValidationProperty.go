package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_NumberValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#max_value TfMemoryStrategy#max_value}.
	// Experimental.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#min_value TfMemoryStrategy#min_value}.
	// Experimental.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
}

