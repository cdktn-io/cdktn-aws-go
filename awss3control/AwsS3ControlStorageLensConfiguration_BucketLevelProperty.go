package awss3control


// Experimental.
type AwsS3ControlStorageLensConfiguration_BucketLevelProperty struct {
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#activity_metrics AwsS3ControlStorageLensConfiguration#activity_metrics}
	// Experimental.
	ActivityMetrics *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// advanced_cost_optimization_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_cost_optimization_metrics AwsS3ControlStorageLensConfiguration#advanced_cost_optimization_metrics}
	// Experimental.
	AdvancedCostOptimizationMetrics *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// advanced_data_protection_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_data_protection_metrics AwsS3ControlStorageLensConfiguration#advanced_data_protection_metrics}
	// Experimental.
	AdvancedDataProtectionMetrics *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// advanced_performance_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#advanced_performance_metrics AwsS3ControlStorageLensConfiguration#advanced_performance_metrics}
	// Experimental.
	AdvancedPerformanceMetrics *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty `field:"optional" json:"advancedPerformanceMetrics" yaml:"advancedPerformanceMetrics"`
	// detailed_status_code_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#detailed_status_code_metrics AwsS3ControlStorageLensConfiguration#detailed_status_code_metrics}
	// Experimental.
	DetailedStatusCodeMetrics *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty `field:"optional" json:"detailedStatusCodeMetrics" yaml:"detailedStatusCodeMetrics"`
	// prefix_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix_level AwsS3ControlStorageLensConfiguration#prefix_level}
	// Experimental.
	PrefixLevel *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

