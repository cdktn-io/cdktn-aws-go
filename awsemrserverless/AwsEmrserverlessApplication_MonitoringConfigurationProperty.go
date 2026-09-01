package awsemrserverless


// Experimental.
type AwsEmrserverlessApplication_MonitoringConfigurationProperty struct {
	// cloudwatch_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#cloudwatch_logging_configuration AwsEmrserverlessApplication#cloudwatch_logging_configuration}
	// Experimental.
	CloudwatchLoggingConfiguration *AwsEmrserverlessApplication_CloudwatchLoggingConfigurationProperty `field:"optional" json:"cloudwatchLoggingConfiguration" yaml:"cloudwatchLoggingConfiguration"`
	// managed_persistence_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#managed_persistence_monitoring_configuration AwsEmrserverlessApplication#managed_persistence_monitoring_configuration}
	// Experimental.
	ManagedPersistenceMonitoringConfiguration *AwsEmrserverlessApplication_ManagedPersistenceMonitoringConfigurationProperty `field:"optional" json:"managedPersistenceMonitoringConfiguration" yaml:"managedPersistenceMonitoringConfiguration"`
	// prometheus_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#prometheus_monitoring_configuration AwsEmrserverlessApplication#prometheus_monitoring_configuration}
	// Experimental.
	PrometheusMonitoringConfiguration *AwsEmrserverlessApplication_PrometheusMonitoringConfigurationProperty `field:"optional" json:"prometheusMonitoringConfiguration" yaml:"prometheusMonitoringConfiguration"`
	// s3_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#s3_monitoring_configuration AwsEmrserverlessApplication#s3_monitoring_configuration}
	// Experimental.
	S3MonitoringConfiguration *AwsEmrserverlessApplication_S3MonitoringConfigurationProperty `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

