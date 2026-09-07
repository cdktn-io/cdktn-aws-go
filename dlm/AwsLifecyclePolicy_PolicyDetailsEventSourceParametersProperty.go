package dlm


// Experimental.
type AwsLifecyclePolicy_PolicyDetailsEventSourceParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#description_regex AwsLifecyclePolicy#description_regex}.
	// Experimental.
	DescriptionRegex *string `field:"required" json:"descriptionRegex" yaml:"descriptionRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#event_type AwsLifecyclePolicy#event_type}.
	// Experimental.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#snapshot_owner AwsLifecyclePolicy#snapshot_owner}.
	// Experimental.
	SnapshotOwner *[]*string `field:"required" json:"snapshotOwner" yaml:"snapshotOwner"`
}

