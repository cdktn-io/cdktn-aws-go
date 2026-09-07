package batch


// Experimental.
type AwsJobDefinition_EksPropertiesProperty struct {
	// pod_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#pod_properties AwsJobDefinition#pod_properties}
	// Experimental.
	PodProperties *AwsJobDefinition_PodPropertiesProperty `field:"required" json:"podProperties" yaml:"podProperties"`
}

