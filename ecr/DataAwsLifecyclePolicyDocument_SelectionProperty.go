package ecr


// Experimental.
type DataAwsLifecyclePolicyDocument_SelectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#count_number DataAwsLifecyclePolicyDocument#count_number}.
	// Experimental.
	CountNumber *float64 `field:"required" json:"countNumber" yaml:"countNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#count_type DataAwsLifecyclePolicyDocument#count_type}.
	// Experimental.
	CountType *string `field:"required" json:"countType" yaml:"countType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#tag_status DataAwsLifecyclePolicyDocument#tag_status}.
	// Experimental.
	TagStatus *string `field:"required" json:"tagStatus" yaml:"tagStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#count_unit DataAwsLifecyclePolicyDocument#count_unit}.
	// Experimental.
	CountUnit *string `field:"optional" json:"countUnit" yaml:"countUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#storage_class DataAwsLifecyclePolicyDocument#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#tag_pattern_list DataAwsLifecyclePolicyDocument#tag_pattern_list}.
	// Experimental.
	TagPatternList *[]*string `field:"optional" json:"tagPatternList" yaml:"tagPatternList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#tag_prefix_list DataAwsLifecyclePolicyDocument#tag_prefix_list}.
	// Experimental.
	TagPrefixList *[]*string `field:"optional" json:"tagPrefixList" yaml:"tagPrefixList"`
}

