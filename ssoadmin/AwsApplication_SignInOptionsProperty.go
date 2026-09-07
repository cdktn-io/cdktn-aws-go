package ssoadmin


// Experimental.
type AwsApplication_SignInOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_application#origin AwsApplication#origin}.
	// Experimental.
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_application#application_url AwsApplication#application_url}.
	// Experimental.
	ApplicationUrl *string `field:"optional" json:"applicationUrl" yaml:"applicationUrl"`
}

