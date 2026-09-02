package awscognitoidp


// Experimental.
type TfUserPool_DeviceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#challenge_required_on_new_device TfUserPool#challenge_required_on_new_device}.
	// Experimental.
	ChallengeRequiredOnNewDevice interface{} `field:"optional" json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#device_only_remembered_on_user_prompt TfUserPool#device_only_remembered_on_user_prompt}.
	// Experimental.
	DeviceOnlyRememberedOnUserPrompt interface{} `field:"optional" json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

