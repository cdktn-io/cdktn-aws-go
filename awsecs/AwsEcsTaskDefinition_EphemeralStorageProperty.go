package awsecs


// Experimental.
type AwsEcsTaskDefinition_EphemeralStorageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#size_in_gib AwsEcsTaskDefinition#size_in_gib}.
	// Experimental.
	SizeInGib *float64 `field:"required" json:"sizeInGib" yaml:"sizeInGib"`
}

