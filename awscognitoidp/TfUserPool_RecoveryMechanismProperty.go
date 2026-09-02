package awscognitoidp


// Experimental.
type TfUserPool_RecoveryMechanismProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name TfUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#priority TfUserPool#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

