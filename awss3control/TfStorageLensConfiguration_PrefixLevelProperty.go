package awss3control


// Experimental.
type TfStorageLensConfiguration_PrefixLevelProperty struct {
	// storage_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_metrics TfStorageLensConfiguration#storage_metrics}
	// Experimental.
	StorageMetrics *TfStorageLensConfiguration_StorageMetricsProperty `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

