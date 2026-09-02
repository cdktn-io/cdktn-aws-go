package awsecs


// Experimental.
type TfDaemonTaskDefinition_RepositoryCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#credentials_parameter TfDaemonTaskDefinition#credentials_parameter}.
	// Experimental.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
}

