package awssagemakerai


// Experimental.
type TfMonitoringSchedule_BatchTransformInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#data_captured_destination_s3_uri TfMonitoringSchedule#data_captured_destination_s3_uri}.
	// Experimental.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// dataset_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#dataset_format TfMonitoringSchedule#dataset_format}
	// Experimental.
	DatasetFormat *TfMonitoringSchedule_DatasetFormatProperty `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#local_path TfMonitoringSchedule#local_path}.
	// Experimental.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#end_time_offset TfMonitoringSchedule#end_time_offset}.
	// Experimental.
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#exclude_features_attribute TfMonitoringSchedule#exclude_features_attribute}.
	// Experimental.
	ExcludeFeaturesAttribute *string `field:"optional" json:"excludeFeaturesAttribute" yaml:"excludeFeaturesAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#features_attribute TfMonitoringSchedule#features_attribute}.
	// Experimental.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#inference_attribute TfMonitoringSchedule#inference_attribute}.
	// Experimental.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#probability_attribute TfMonitoringSchedule#probability_attribute}.
	// Experimental.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#probability_threshold_attribute TfMonitoringSchedule#probability_threshold_attribute}.
	// Experimental.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_data_distribution_type TfMonitoringSchedule#s3_data_distribution_type}.
	// Experimental.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_input_mode TfMonitoringSchedule#s3_input_mode}.
	// Experimental.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#start_time_offset TfMonitoringSchedule#start_time_offset}.
	// Experimental.
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

