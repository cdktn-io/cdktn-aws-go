package awsdynamodb


// Experimental.
type TfTable_ReplicaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#region_name TfTable#region_name}.
	// Experimental.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#consistency_mode TfTable#consistency_mode}.
	// Experimental.
	ConsistencyMode *string `field:"optional" json:"consistencyMode" yaml:"consistencyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#deletion_protection_enabled TfTable#deletion_protection_enabled}.
	// Experimental.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#kms_key_arn TfTable#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#point_in_time_recovery TfTable#point_in_time_recovery}.
	// Experimental.
	PointInTimeRecovery interface{} `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#propagate_tags TfTable#propagate_tags}.
	// Experimental.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

