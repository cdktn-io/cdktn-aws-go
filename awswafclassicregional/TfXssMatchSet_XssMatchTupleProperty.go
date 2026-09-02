package awswafclassicregional


// Experimental.
type TfXssMatchSet_XssMatchTupleProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_xss_match_set#field_to_match TfXssMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *TfXssMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_xss_match_set#text_transformation TfXssMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

