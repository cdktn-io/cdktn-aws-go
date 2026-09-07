package cognitoidp


// Experimental.
type AwsUserPool_WebAuthnConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#relying_party_id AwsUserPool#relying_party_id}.
	// Experimental.
	RelyingPartyId *string `field:"optional" json:"relyingPartyId" yaml:"relyingPartyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#user_verification AwsUserPool#user_verification}.
	// Experimental.
	UserVerification *string `field:"optional" json:"userVerification" yaml:"userVerification"`
}

