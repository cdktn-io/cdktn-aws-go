package quicksight


// Experimental.
type AwsKeyRegistration_KeyRegistrationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_key_registration#key_arn AwsKeyRegistration#key_arn}.
	// Experimental.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_key_registration#default_key AwsKeyRegistration#default_key}.
	// Experimental.
	DefaultKey interface{} `field:"optional" json:"defaultKey" yaml:"defaultKey"`
}

