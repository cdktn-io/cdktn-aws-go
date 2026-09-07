package cognitoidp


// Experimental.
type AwsUserPool_PasswordPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#minimum_length AwsUserPool#minimum_length}.
	// Experimental.
	MinimumLength *float64 `field:"optional" json:"minimumLength" yaml:"minimumLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#password_history_size AwsUserPool#password_history_size}.
	// Experimental.
	PasswordHistorySize *float64 `field:"optional" json:"passwordHistorySize" yaml:"passwordHistorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#require_lowercase AwsUserPool#require_lowercase}.
	// Experimental.
	RequireLowercase interface{} `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#require_numbers AwsUserPool#require_numbers}.
	// Experimental.
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#require_symbols AwsUserPool#require_symbols}.
	// Experimental.
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#require_uppercase AwsUserPool#require_uppercase}.
	// Experimental.
	RequireUppercase interface{} `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#temporary_password_validity_days AwsUserPool#temporary_password_validity_days}.
	// Experimental.
	TemporaryPasswordValidityDays *float64 `field:"optional" json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

