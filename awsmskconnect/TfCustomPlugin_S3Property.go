package awsmskconnect


// Experimental.
type TfCustomPlugin_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_custom_plugin#bucket_arn TfCustomPlugin#bucket_arn}.
	// Experimental.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_custom_plugin#file_key TfCustomPlugin#file_key}.
	// Experimental.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_custom_plugin#object_version TfCustomPlugin#object_version}.
	// Experimental.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

