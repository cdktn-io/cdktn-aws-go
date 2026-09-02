package awswafclassic


// Experimental.
type TfRegexMatchSet_RegexMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#field_to_match TfRegexMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *TfRegexMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#regex_pattern_set_id TfRegexMatchSet#regex_pattern_set_id}.
	// Experimental.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#text_transformation TfRegexMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

