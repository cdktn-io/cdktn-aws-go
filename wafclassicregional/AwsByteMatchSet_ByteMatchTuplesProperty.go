package wafclassicregional


// Experimental.
type AwsByteMatchSet_ByteMatchTuplesProperty struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_byte_match_set#field_to_match AwsByteMatchSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsByteMatchSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_byte_match_set#positional_constraint AwsByteMatchSet#positional_constraint}.
	// Experimental.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_byte_match_set#text_transformation AwsByteMatchSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_byte_match_set#target_string AwsByteMatchSet#target_string}.
	// Experimental.
	TargetString *string `field:"optional" json:"targetString" yaml:"targetString"`
}

