package awswafclassic


// Experimental.
type AwsWafXssMatchSet_XssMatchTuplesProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_xss_match_set#field_to_match AwsWafXssMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafXssMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_xss_match_set#text_transformation AwsWafXssMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

