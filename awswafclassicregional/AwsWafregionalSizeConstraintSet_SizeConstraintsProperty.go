package awswafclassicregional


// Experimental.
type AwsWafregionalSizeConstraintSet_SizeConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_size_constraint_set#comparison_operator AwsWafregionalSizeConstraintSet#comparison_operator}.
	// Experimental.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_size_constraint_set#field_to_match AwsWafregionalSizeConstraintSet#field_to_match}
	// Experimental.
	FieldToMatch *AwsWafregionalSizeConstraintSet_FieldToMatchProperty `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_size_constraint_set#size AwsWafregionalSizeConstraintSet#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_size_constraint_set#text_transformation AwsWafregionalSizeConstraintSet#text_transformation}.
	// Experimental.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

