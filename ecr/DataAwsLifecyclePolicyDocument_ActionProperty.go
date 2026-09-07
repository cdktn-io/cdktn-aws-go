package ecr


// Experimental.
type DataAwsLifecyclePolicyDocument_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#type DataAwsLifecyclePolicyDocument#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#target_storage_class DataAwsLifecyclePolicyDocument#target_storage_class}.
	// Experimental.
	TargetStorageClass *string `field:"optional" json:"targetStorageClass" yaml:"targetStorageClass"`
}

