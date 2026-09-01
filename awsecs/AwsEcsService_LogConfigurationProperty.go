package awsecs


// Experimental.
type AwsEcsService_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#log_driver AwsEcsService#log_driver}.
	// Experimental.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#options AwsEcsService#options}.
	// Experimental.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#secret_option AwsEcsService#secret_option}
	// Experimental.
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

