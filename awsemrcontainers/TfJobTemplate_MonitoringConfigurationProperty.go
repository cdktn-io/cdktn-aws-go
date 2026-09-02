package awsemrcontainers


// Experimental.
type TfJobTemplate_MonitoringConfigurationProperty struct {
	// cloud_watch_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#cloud_watch_monitoring_configuration TfJobTemplate#cloud_watch_monitoring_configuration}
	// Experimental.
	CloudWatchMonitoringConfiguration *TfJobTemplate_CloudWatchMonitoringConfigurationProperty `field:"optional" json:"cloudWatchMonitoringConfiguration" yaml:"cloudWatchMonitoringConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#persistent_app_ui TfJobTemplate#persistent_app_ui}.
	// Experimental.
	PersistentAppUi *string `field:"optional" json:"persistentAppUi" yaml:"persistentAppUi"`
	// s3_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#s3_monitoring_configuration TfJobTemplate#s3_monitoring_configuration}
	// Experimental.
	S3MonitoringConfiguration *TfJobTemplate_S3MonitoringConfigurationProperty `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

