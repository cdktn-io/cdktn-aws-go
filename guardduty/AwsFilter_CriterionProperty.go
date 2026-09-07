package guardduty


// Experimental.
type AwsFilter_CriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#field AwsFilter#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#equals AwsFilter#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#greater_than AwsFilter#greater_than}.
	// Experimental.
	GreaterThan *string `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#greater_than_or_equal AwsFilter#greater_than_or_equal}.
	// Experimental.
	GreaterThanOrEqual *string `field:"optional" json:"greaterThanOrEqual" yaml:"greaterThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#less_than AwsFilter#less_than}.
	// Experimental.
	LessThan *string `field:"optional" json:"lessThan" yaml:"lessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#less_than_or_equal AwsFilter#less_than_or_equal}.
	// Experimental.
	LessThanOrEqual *string `field:"optional" json:"lessThanOrEqual" yaml:"lessThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#matches AwsFilter#matches}.
	// Experimental.
	Matches *[]*string `field:"optional" json:"matches" yaml:"matches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#not_equals AwsFilter#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#not_matches AwsFilter#not_matches}.
	// Experimental.
	NotMatches *[]*string `field:"optional" json:"notMatches" yaml:"notMatches"`
}

