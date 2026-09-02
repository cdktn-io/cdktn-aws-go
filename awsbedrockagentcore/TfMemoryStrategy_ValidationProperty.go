package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_ValidationProperty struct {
	// number_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#number_validation TfMemoryStrategy#number_validation}
	// Experimental.
	NumberValidation interface{} `field:"optional" json:"numberValidation" yaml:"numberValidation"`
	// string_list_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#string_list_validation TfMemoryStrategy#string_list_validation}
	// Experimental.
	StringListValidation interface{} `field:"optional" json:"stringListValidation" yaml:"stringListValidation"`
	// string_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#string_validation TfMemoryStrategy#string_validation}
	// Experimental.
	StringValidation interface{} `field:"optional" json:"stringValidation" yaml:"stringValidation"`
}

