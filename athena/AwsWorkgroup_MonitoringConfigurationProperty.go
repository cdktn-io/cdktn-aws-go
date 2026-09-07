package athena


// Experimental.
type AwsWorkgroup_MonitoringConfigurationProperty struct {
	// cloud_watch_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#cloud_watch_logging_configuration AwsWorkgroup#cloud_watch_logging_configuration}
	// Experimental.
	CloudWatchLoggingConfiguration *AwsWorkgroup_CloudWatchLoggingConfigurationProperty `field:"optional" json:"cloudWatchLoggingConfiguration" yaml:"cloudWatchLoggingConfiguration"`
	// managed_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#managed_logging_configuration AwsWorkgroup#managed_logging_configuration}
	// Experimental.
	ManagedLoggingConfiguration *AwsWorkgroup_ManagedLoggingConfigurationProperty `field:"optional" json:"managedLoggingConfiguration" yaml:"managedLoggingConfiguration"`
	// s3_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#s3_logging_configuration AwsWorkgroup#s3_logging_configuration}
	// Experimental.
	S3LoggingConfiguration *AwsWorkgroup_S3LoggingConfigurationProperty `field:"optional" json:"s3LoggingConfiguration" yaml:"s3LoggingConfiguration"`
}

