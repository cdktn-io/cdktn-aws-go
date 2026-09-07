package cognitoidp


// Experimental.
type AwsUserPool_NumberAttributeConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#max_value AwsUserPool#max_value}.
	// Experimental.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#min_value AwsUserPool#min_value}.
	// Experimental.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

