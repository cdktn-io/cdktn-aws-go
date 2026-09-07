package backup


// Experimental.
type AwsSelection_ConditionProperty struct {
	// string_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#string_equals AwsSelection#string_equals}
	// Experimental.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// string_like block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#string_like AwsSelection#string_like}
	// Experimental.
	StringLike interface{} `field:"optional" json:"stringLike" yaml:"stringLike"`
	// string_not_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#string_not_equals AwsSelection#string_not_equals}
	// Experimental.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// string_not_like block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_selection#string_not_like AwsSelection#string_not_like}
	// Experimental.
	StringNotLike interface{} `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}

