package awscognitoidp


// Experimental.
type AwsCognitoUserPool_RecoveryMechanismProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name AwsCognitoUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#priority AwsCognitoUserPool#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

