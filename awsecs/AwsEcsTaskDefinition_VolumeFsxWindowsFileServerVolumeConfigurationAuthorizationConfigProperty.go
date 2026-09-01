package awsecs


// Experimental.
type AwsEcsTaskDefinition_VolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#credentials_parameter AwsEcsTaskDefinition#credentials_parameter}.
	// Experimental.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#domain AwsEcsTaskDefinition#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

