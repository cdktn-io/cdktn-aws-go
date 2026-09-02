package awsecs


// Experimental.
type TfTaskDefinition_EphemeralStorageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#size_in_gib TfTaskDefinition#size_in_gib}.
	// Experimental.
	SizeInGib *float64 `field:"required" json:"sizeInGib" yaml:"sizeInGib"`
}

