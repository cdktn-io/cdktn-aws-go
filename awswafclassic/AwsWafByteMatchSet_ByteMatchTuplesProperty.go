package awswafclassic


// Experimental.
type AwsWafByteMatchSet_ByteMatchTuplesProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_byte_match_set#field_to_match AwsWafByteMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafByteMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_byte_match_set#positional_constraint AwsWafByteMatchSet#positional_constraint}.
	// Experimental.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_byte_match_set#text_transformation AwsWafByteMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_byte_match_set#target_string AwsWafByteMatchSet#target_string}.
	// Experimental.
	TargetString *string `field:"optional" json:"targetString" yaml:"targetString"`
}

