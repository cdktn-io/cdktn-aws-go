package awsbatch


// Experimental.
type TfJobDefinition_EksPropertiesProperty struct {
	// pod_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#pod_properties TfJobDefinition#pod_properties}
	// Experimental.
	PodProperties *TfJobDefinition_PodPropertiesProperty `field:"required" json:"podProperties" yaml:"podProperties"`
}

