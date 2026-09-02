package awseventbridgepipes


// Experimental.
type TfPipe_OverridesProperty struct {
	// container_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#container_override TfPipe#container_override}
	// Experimental.
	ContainerOverride interface{} `field:"optional" json:"containerOverride" yaml:"containerOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cpu TfPipe#cpu}.
	// Experimental.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#ephemeral_storage TfPipe#ephemeral_storage}
	// Experimental.
	EphemeralStorage *TfPipe_EphemeralStorageProperty `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#execution_role_arn TfPipe#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// inference_accelerator_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#inference_accelerator_override TfPipe#inference_accelerator_override}
	// Experimental.
	InferenceAcceleratorOverride interface{} `field:"optional" json:"inferenceAcceleratorOverride" yaml:"inferenceAcceleratorOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#memory TfPipe#memory}.
	// Experimental.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#task_role_arn TfPipe#task_role_arn}.
	// Experimental.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

