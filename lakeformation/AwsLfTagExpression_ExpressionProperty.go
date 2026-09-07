package lakeformation


// Experimental.
type AwsLfTagExpression_ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_lf_tag_expression#tag_key AwsLfTagExpression#tag_key}.
	// Experimental.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_lf_tag_expression#tag_values AwsLfTagExpression#tag_values}.
	// Experimental.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

