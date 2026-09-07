package elementalmedialive


// Experimental.
type AwsInput_SourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#password_param AwsInput#password_param}.
	// Experimental.
	PasswordParam *string `field:"required" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#url AwsInput#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#username AwsInput#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

