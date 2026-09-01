package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_AudioOnlyImageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#uri AwsMedialiveChannel#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#password_param AwsMedialiveChannel#password_param}.
	// Experimental.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#username AwsMedialiveChannel#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

