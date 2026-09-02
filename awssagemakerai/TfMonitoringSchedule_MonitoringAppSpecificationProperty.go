package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringAppSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#image_uri TfMonitoringSchedule#image_uri}.
	// Experimental.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#container_arguments TfMonitoringSchedule#container_arguments}.
	// Experimental.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#container_entrypoint TfMonitoringSchedule#container_entrypoint}.
	// Experimental.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#post_analytics_processor_source_uri TfMonitoringSchedule#post_analytics_processor_source_uri}.
	// Experimental.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#record_preprocessor_source_uri TfMonitoringSchedule#record_preprocessor_source_uri}.
	// Experimental.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

