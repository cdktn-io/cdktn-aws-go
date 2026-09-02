package awslakeformation


// Experimental.
type TfLfTagExpression_ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_lf_tag_expression#tag_key TfLfTagExpression#tag_key}.
	// Experimental.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_lf_tag_expression#tag_values TfLfTagExpression#tag_values}.
	// Experimental.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

