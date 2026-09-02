package awskinesisanalyticsv2


// Experimental.
type TfApplication_MonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#configuration_type TfApplication#configuration_type}.
	// Experimental.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#log_level TfApplication#log_level}.
	// Experimental.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#metrics_level TfApplication#metrics_level}.
	// Experimental.
	MetricsLevel *string `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
}

