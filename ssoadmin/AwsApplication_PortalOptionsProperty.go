package ssoadmin


// Experimental.
type AwsApplication_PortalOptionsProperty struct {
	// sign_in_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_application#sign_in_options AwsApplication#sign_in_options}
	// Experimental.
	SignInOptions interface{} `field:"optional" json:"signInOptions" yaml:"signInOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_application#visibility AwsApplication#visibility}.
	// Experimental.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

