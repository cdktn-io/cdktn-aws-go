package awsecs


// Experimental.
type AwsEcsDaemonTaskDefinition_RepositoryCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#credentials_parameter AwsEcsDaemonTaskDefinition#credentials_parameter}.
	// Experimental.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
}

