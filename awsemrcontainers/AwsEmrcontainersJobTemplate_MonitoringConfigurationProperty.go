package awsemrcontainers


// Experimental.
type AwsEmrcontainersJobTemplate_MonitoringConfigurationProperty struct {
	// cloud_watch_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#cloud_watch_monitoring_configuration AwsEmrcontainersJobTemplate#cloud_watch_monitoring_configuration}
	// Experimental.
	CloudWatchMonitoringConfiguration *AwsEmrcontainersJobTemplate_CloudWatchMonitoringConfigurationProperty `field:"optional" json:"cloudWatchMonitoringConfiguration" yaml:"cloudWatchMonitoringConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#persistent_app_ui AwsEmrcontainersJobTemplate#persistent_app_ui}.
	// Experimental.
	PersistentAppUi *string `field:"optional" json:"persistentAppUi" yaml:"persistentAppUi"`
	// s3_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#s3_monitoring_configuration AwsEmrcontainersJobTemplate#s3_monitoring_configuration}
	// Experimental.
	S3MonitoringConfiguration *AwsEmrcontainersJobTemplate_S3MonitoringConfigurationProperty `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

