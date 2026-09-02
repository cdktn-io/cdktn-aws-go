package awsecs


// Experimental.
type TfTaskDefinition_DockerVolumeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#autoprovision TfTaskDefinition#autoprovision}.
	// Experimental.
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#driver TfTaskDefinition#driver}.
	// Experimental.
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#driver_opts TfTaskDefinition#driver_opts}.
	// Experimental.
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#labels TfTaskDefinition#labels}.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#scope TfTaskDefinition#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

