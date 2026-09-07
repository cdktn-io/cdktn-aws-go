package cognitoidp


// Experimental.
type AwsUserPool_StringAttributeConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#max_length AwsUserPool#max_length}.
	// Experimental.
	MaxLength *string `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#min_length AwsUserPool#min_length}.
	// Experimental.
	MinLength *string `field:"optional" json:"minLength" yaml:"minLength"`
}

