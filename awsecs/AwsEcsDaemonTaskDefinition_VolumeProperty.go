package awsecs


// Experimental.
type AwsEcsDaemonTaskDefinition_VolumeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name AwsEcsDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#host AwsEcsDaemonTaskDefinition#host}
	// Experimental.
	Host interface{} `field:"optional" json:"host" yaml:"host"`
}

