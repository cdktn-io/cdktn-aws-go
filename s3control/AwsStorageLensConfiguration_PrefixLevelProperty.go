package s3control


// Experimental.
type AwsStorageLensConfiguration_PrefixLevelProperty struct {
	// storage_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_metrics AwsStorageLensConfiguration#storage_metrics}
	// Experimental.
	StorageMetrics *AwsStorageLensConfiguration_StorageMetricsProperty `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

