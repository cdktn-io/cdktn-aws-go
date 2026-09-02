package awsecs


// Experimental.
type TfTaskDefinition_ProxyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#container_name TfTaskDefinition#container_name}.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#properties TfTaskDefinition#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#type TfTaskDefinition#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

