package awswafclassic


// Experimental.
type AwsWafRegexMatchSet_RegexMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#field_to_match AwsWafRegexMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafRegexMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#regex_pattern_set_id AwsWafRegexMatchSet#regex_pattern_set_id}.
	// Experimental.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_regex_match_set#text_transformation AwsWafRegexMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

