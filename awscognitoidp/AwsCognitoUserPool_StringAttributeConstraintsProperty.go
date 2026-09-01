package awscognitoidp


// Experimental.
type AwsCognitoUserPool_StringAttributeConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#max_length AwsCognitoUserPool#max_length}.
	// Experimental.
	MaxLength *string `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#min_length AwsCognitoUserPool#min_length}.
	// Experimental.
	MinLength *string `field:"optional" json:"minLength" yaml:"minLength"`
}

