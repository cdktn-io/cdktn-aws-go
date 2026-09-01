package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_KendraConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#kendra_index AwsLexv2ModelsIntent#kendra_index}.
	// Experimental.
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#query_filter_string AwsLexv2ModelsIntent#query_filter_string}.
	// Experimental.
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#query_filter_string_enabled AwsLexv2ModelsIntent#query_filter_string_enabled}.
	// Experimental.
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

