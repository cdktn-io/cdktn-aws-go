package cognitoidp


// Experimental.
type AwsUserPool_RecoveryMechanismProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name AwsUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#priority AwsUserPool#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

