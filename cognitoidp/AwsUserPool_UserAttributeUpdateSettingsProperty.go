package cognitoidp


// Experimental.
type AwsUserPool_UserAttributeUpdateSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#attributes_require_verification_before_update AwsUserPool#attributes_require_verification_before_update}.
	// Experimental.
	AttributesRequireVerificationBeforeUpdate *[]*string `field:"required" json:"attributesRequireVerificationBeforeUpdate" yaml:"attributesRequireVerificationBeforeUpdate"`
}

