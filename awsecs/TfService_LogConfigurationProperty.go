package awsecs


// Experimental.
type TfService_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#log_driver TfService#log_driver}.
	// Experimental.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#options TfService#options}.
	// Experimental.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#secret_option TfService#secret_option}
	// Experimental.
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

