package awscognitoidp


// Experimental.
type TfManagedUserPoolClient_RefreshTokenRotationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#feature TfManagedUserPoolClient#feature}.
	// Experimental.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#retry_grace_period_seconds TfManagedUserPoolClient#retry_grace_period_seconds}.
	// Experimental.
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

