package awss3control


// Experimental.
type AwsS3ControlStorageLensConfiguration_PrefixLevelProperty struct {
	// storage_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_metrics AwsS3ControlStorageLensConfiguration#storage_metrics}
	// Experimental.
	StorageMetrics *AwsS3ControlStorageLensConfiguration_StorageMetricsProperty `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

