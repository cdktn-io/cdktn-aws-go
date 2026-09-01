package awswafclassicregional


// Experimental.
type AwsWafregionalSqlInjectionMatchSet_SqlInjectionMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_sql_injection_match_set#field_to_match AwsWafregionalSqlInjectionMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafregionalSqlInjectionMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_sql_injection_match_set#text_transformation AwsWafregionalSqlInjectionMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

