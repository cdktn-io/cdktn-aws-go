package cognitoidp


// Experimental.
type AwsUserPoolClient_RefreshTokenRotationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#feature AwsUserPoolClient#feature}.
	// Experimental.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#retry_grace_period_seconds AwsUserPoolClient#retry_grace_period_seconds}.
	// Experimental.
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

