package awsdlm


// Experimental.
type TfLifecyclePolicy_PolicyDetailsProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#action TfLifecyclePolicy#action}
	// Experimental.
	Action *TfLifecyclePolicy_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#copy_tags TfLifecyclePolicy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#create_interval TfLifecyclePolicy#create_interval}.
	// Experimental.
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#event_source TfLifecyclePolicy#event_source}
	// Experimental.
	EventSource *TfLifecyclePolicy_EventSourceProperty `field:"optional" json:"eventSource" yaml:"eventSource"`
	// exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclusions TfLifecyclePolicy#exclusions}
	// Experimental.
	Exclusions *TfLifecyclePolicy_ExclusionsProperty `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#extend_deletion TfLifecyclePolicy#extend_deletion}.
	// Experimental.
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#parameters TfLifecyclePolicy#parameters}
	// Experimental.
	Parameters *TfLifecyclePolicy_PolicyDetailsParametersProperty `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#policy_language TfLifecyclePolicy#policy_language}.
	// Experimental.
	PolicyLanguage *string `field:"optional" json:"policyLanguage" yaml:"policyLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#policy_type TfLifecyclePolicy#policy_type}.
	// Experimental.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_locations TfLifecyclePolicy#resource_locations}.
	// Experimental.
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_type TfLifecyclePolicy#resource_type}.
	// Experimental.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_types TfLifecyclePolicy#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_interval TfLifecyclePolicy#retain_interval}.
	// Experimental.
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#schedule TfLifecyclePolicy#schedule}
	// Experimental.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target_tags TfLifecyclePolicy#target_tags}.
	// Experimental.
	TargetTags *map[string]*string `field:"optional" json:"targetTags" yaml:"targetTags"`
}

