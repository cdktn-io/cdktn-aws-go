package awsguardduty


// Experimental.
type AwsGuarddutyFilter_CriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#field AwsGuarddutyFilter#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#equals AwsGuarddutyFilter#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#greater_than AwsGuarddutyFilter#greater_than}.
	// Experimental.
	GreaterThan *string `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#greater_than_or_equal AwsGuarddutyFilter#greater_than_or_equal}.
	// Experimental.
	GreaterThanOrEqual *string `field:"optional" json:"greaterThanOrEqual" yaml:"greaterThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#less_than AwsGuarddutyFilter#less_than}.
	// Experimental.
	LessThan *string `field:"optional" json:"lessThan" yaml:"lessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#less_than_or_equal AwsGuarddutyFilter#less_than_or_equal}.
	// Experimental.
	LessThanOrEqual *string `field:"optional" json:"lessThanOrEqual" yaml:"lessThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#matches AwsGuarddutyFilter#matches}.
	// Experimental.
	Matches *[]*string `field:"optional" json:"matches" yaml:"matches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#not_equals AwsGuarddutyFilter#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#not_matches AwsGuarddutyFilter#not_matches}.
	// Experimental.
	NotMatches *[]*string `field:"optional" json:"notMatches" yaml:"notMatches"`
}

