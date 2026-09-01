package awswafclassicregional


// Experimental.
type AwsWafregionalRegexMatchSet_RegexMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_regex_match_set#field_to_match AwsWafregionalRegexMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafregionalRegexMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_regex_match_set#regex_pattern_set_id AwsWafregionalRegexMatchSet#regex_pattern_set_id}.
	// Experimental.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_regex_match_set#text_transformation AwsWafregionalRegexMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

