package s3control


// Experimental.
type AwsStorageLensConfiguration_AccountLevelProperty struct {
	// bucket_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#bucket_level AwsStorageLensConfiguration#bucket_level}
	// Experimental.
	BucketLevel *AwsStorageLensConfiguration_BucketLevelProperty `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#activity_metrics AwsStorageLensConfiguration#activity_metrics}
	// Experimental.
	ActivityMetrics *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// advanced_cost_optimization_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_cost_optimization_metrics AwsStorageLensConfiguration#advanced_cost_optimization_metrics}
	// Experimental.
	AdvancedCostOptimizationMetrics *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// advanced_data_protection_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_data_protection_metrics AwsStorageLensConfiguration#advanced_data_protection_metrics}
	// Experimental.
	AdvancedDataProtectionMetrics *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// advanced_performance_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_performance_metrics AwsStorageLensConfiguration#advanced_performance_metrics}
	// Experimental.
	AdvancedPerformanceMetrics *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty `field:"optional" json:"advancedPerformanceMetrics" yaml:"advancedPerformanceMetrics"`
	// detailed_status_code_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#detailed_status_code_metrics AwsStorageLensConfiguration#detailed_status_code_metrics}
	// Experimental.
	DetailedStatusCodeMetrics *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty `field:"optional" json:"detailedStatusCodeMetrics" yaml:"detailedStatusCodeMetrics"`
}

