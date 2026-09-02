package awsecr


// Experimental.
type DataTfLifecyclePolicyDocument_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#type DataTfLifecyclePolicyDocument#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#target_storage_class DataTfLifecyclePolicyDocument#target_storage_class}.
	// Experimental.
	TargetStorageClass *string `field:"optional" json:"targetStorageClass" yaml:"targetStorageClass"`
}

