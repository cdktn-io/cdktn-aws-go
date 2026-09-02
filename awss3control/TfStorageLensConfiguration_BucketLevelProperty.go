package awss3control


// Experimental.
type TfStorageLensConfiguration_BucketLevelProperty struct {
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#activity_metrics TfStorageLensConfiguration#activity_metrics}
	// Experimental.
	ActivityMetrics *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// advanced_cost_optimization_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_cost_optimization_metrics TfStorageLensConfiguration#advanced_cost_optimization_metrics}
	// Experimental.
	AdvancedCostOptimizationMetrics *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// advanced_data_protection_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_data_protection_metrics TfStorageLensConfiguration#advanced_data_protection_metrics}
	// Experimental.
	AdvancedDataProtectionMetrics *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// advanced_performance_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_performance_metrics TfStorageLensConfiguration#advanced_performance_metrics}
	// Experimental.
	AdvancedPerformanceMetrics *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty `field:"optional" json:"advancedPerformanceMetrics" yaml:"advancedPerformanceMetrics"`
	// detailed_status_code_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#detailed_status_code_metrics TfStorageLensConfiguration#detailed_status_code_metrics}
	// Experimental.
	DetailedStatusCodeMetrics *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty `field:"optional" json:"detailedStatusCodeMetrics" yaml:"detailedStatusCodeMetrics"`
	// prefix_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix_level TfStorageLensConfiguration#prefix_level}
	// Experimental.
	PrefixLevel *TfStorageLensConfiguration_PrefixLevelProperty `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

