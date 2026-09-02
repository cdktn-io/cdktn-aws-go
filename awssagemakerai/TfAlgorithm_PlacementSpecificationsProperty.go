package awssagemakerai


// Experimental.
type TfAlgorithm_PlacementSpecificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_count TfAlgorithm#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#ultra_server_id TfAlgorithm#ultra_server_id}.
	// Experimental.
	UltraServerId *string `field:"optional" json:"ultraServerId" yaml:"ultraServerId"`
}

