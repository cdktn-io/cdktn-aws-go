package wafclassic


// Experimental.
type AwsXssMatchSet_XssMatchTuplesProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_xss_match_set#field_to_match AwsXssMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsXssMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_xss_match_set#text_transformation AwsXssMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

