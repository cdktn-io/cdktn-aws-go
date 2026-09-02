package awsecs


// Experimental.
type TfService_ManagedEbsVolumeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn TfService#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#encrypted TfService#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#file_system_type TfService#file_system_type}.
	// Experimental.
	FileSystemType *string `field:"optional" json:"fileSystemType" yaml:"fileSystemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#iops TfService#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#kms_key_id TfService#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#size_in_gb TfService#size_in_gb}.
	// Experimental.
	SizeInGb *float64 `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#snapshot_id TfService#snapshot_id}.
	// Experimental.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// tag_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#tag_specifications TfService#tag_specifications}
	// Experimental.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#throughput TfService#throughput}.
	// Experimental.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#volume_initialization_rate TfService#volume_initialization_rate}.
	// Experimental.
	VolumeInitializationRate *float64 `field:"optional" json:"volumeInitializationRate" yaml:"volumeInitializationRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#volume_type TfService#volume_type}.
	// Experimental.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

