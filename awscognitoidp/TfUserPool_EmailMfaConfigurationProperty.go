package awscognitoidp


// Experimental.
type TfUserPool_EmailMfaConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#message TfUserPool#message}.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#subject TfUserPool#subject}.
	// Experimental.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

