package awsbatch


// Experimental.
type AwsBatchJobDefinition_EksPropertiesProperty struct {
	// pod_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#pod_properties AwsBatchJobDefinition#pod_properties}
	// Experimental.
	PodProperties *AwsBatchJobDefinition_PodPropertiesProperty `field:"required" json:"podProperties" yaml:"podProperties"`
}

