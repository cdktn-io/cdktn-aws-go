package awselementalmedialive


// Experimental.
type AwsMedialiveInput_SourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#password_param AwsMedialiveInput#password_param}.
	// Experimental.
	PasswordParam *string `field:"required" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#url AwsMedialiveInput#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_input#username AwsMedialiveInput#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

