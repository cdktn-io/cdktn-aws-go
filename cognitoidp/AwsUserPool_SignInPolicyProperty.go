package cognitoidp


// Experimental.
type AwsUserPool_SignInPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#allowed_first_auth_factors AwsUserPool#allowed_first_auth_factors}.
	// Experimental.
	AllowedFirstAuthFactors *[]*string `field:"optional" json:"allowedFirstAuthFactors" yaml:"allowedFirstAuthFactors"`
}

