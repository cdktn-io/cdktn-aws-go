package awslexv2models


// Experimental.
type TfSlotType_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#kms_key_arn TfSlotType#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#s3_bucket_name TfSlotType#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#s3_object_key TfSlotType#s3_object_key}.
	// Experimental.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

