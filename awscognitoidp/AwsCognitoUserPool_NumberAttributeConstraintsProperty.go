package awscognitoidp


// Experimental.
type AwsCognitoUserPool_NumberAttributeConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#max_value AwsCognitoUserPool#max_value}.
	// Experimental.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#min_value AwsCognitoUserPool#min_value}.
	// Experimental.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

