package awswafclassicregional


// Experimental.
type AwsWafregionalXssMatchSet_XssMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_xss_match_set#field_to_match AwsWafregionalXssMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafregionalXssMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_xss_match_set#text_transformation AwsWafregionalXssMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

