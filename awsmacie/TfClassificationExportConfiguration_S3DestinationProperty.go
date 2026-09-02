package awsmacie


// Experimental.
type TfClassificationExportConfiguration_S3DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_export_configuration#bucket_name TfClassificationExportConfiguration#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_export_configuration#kms_key_arn TfClassificationExportConfiguration#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_export_configuration#key_prefix TfClassificationExportConfiguration#key_prefix}.
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

