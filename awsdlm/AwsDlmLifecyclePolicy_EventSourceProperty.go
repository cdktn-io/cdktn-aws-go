package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_EventSourceProperty struct {
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#parameters AwsDlmLifecyclePolicy#parameters}
	// Experimental.
	Parameters *AwsDlmLifecyclePolicy_PolicyDetailsEventSourceParametersProperty `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#type AwsDlmLifecyclePolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

