package awswafclassic


// Experimental.
type TfSqlInjectionMatchSet_SqlInjectionMatchTuplesProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_sql_injection_match_set#field_to_match TfSqlInjectionMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *TfSqlInjectionMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_sql_injection_match_set#text_transformation TfSqlInjectionMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

