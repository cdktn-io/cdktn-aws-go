package awsdlm


// Experimental.
type TfLifecyclePolicy_ActionProperty struct {
	// cross_region_copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cross_region_copy TfLifecyclePolicy#cross_region_copy}
	// Experimental.
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#name TfLifecyclePolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

