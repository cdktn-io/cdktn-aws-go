package ecs


// Experimental.
type AwsTaskDefinition_VolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#credentials_parameter AwsTaskDefinition#credentials_parameter}.
	// Experimental.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#domain AwsTaskDefinition#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

