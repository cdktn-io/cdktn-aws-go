package dlm


// Experimental.
type AwsLifecyclePolicy_PolicyDetailsProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#action AwsLifecyclePolicy#action}
	// Experimental.
	Action *AwsLifecyclePolicy_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#copy_tags AwsLifecyclePolicy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#create_interval AwsLifecyclePolicy#create_interval}.
	// Experimental.
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#event_source AwsLifecyclePolicy#event_source}
	// Experimental.
	EventSource *AwsLifecyclePolicy_EventSourceProperty `field:"optional" json:"eventSource" yaml:"eventSource"`
	// exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclusions AwsLifecyclePolicy#exclusions}
	// Experimental.
	Exclusions *AwsLifecyclePolicy_ExclusionsProperty `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#extend_deletion AwsLifecyclePolicy#extend_deletion}.
	// Experimental.
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#parameters AwsLifecyclePolicy#parameters}
	// Experimental.
	Parameters *AwsLifecyclePolicy_PolicyDetailsParametersProperty `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#policy_language AwsLifecyclePolicy#policy_language}.
	// Experimental.
	PolicyLanguage *string `field:"optional" json:"policyLanguage" yaml:"policyLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#policy_type AwsLifecyclePolicy#policy_type}.
	// Experimental.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_locations AwsLifecyclePolicy#resource_locations}.
	// Experimental.
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_type AwsLifecyclePolicy#resource_type}.
	// Experimental.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#resource_types AwsLifecyclePolicy#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_interval AwsLifecyclePolicy#retain_interval}.
	// Experimental.
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#schedule AwsLifecyclePolicy#schedule}
	// Experimental.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target_tags AwsLifecyclePolicy#target_tags}.
	// Experimental.
	TargetTags *map[string]*string `field:"optional" json:"targetTags" yaml:"targetTags"`
}

