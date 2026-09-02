package awslexv2models


// Experimental.
type TfIntent_QnaIntentConfigurationDataSourceConfigurationKendraConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#kendra_index TfIntent#kendra_index}.
	// Experimental.
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response TfIntent#exact_response}.
	// Experimental.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#query_filter_string TfIntent#query_filter_string}.
	// Experimental.
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#query_filter_string_enabled TfIntent#query_filter_string_enabled}.
	// Experimental.
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

