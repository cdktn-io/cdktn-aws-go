package awsathena


// Experimental.
type AwsAthenaWorkgroup_MonitoringConfigurationProperty struct {
	// cloud_watch_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#cloud_watch_logging_configuration AwsAthenaWorkgroup#cloud_watch_logging_configuration}
	// Experimental.
	CloudWatchLoggingConfiguration *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty `field:"optional" json:"cloudWatchLoggingConfiguration" yaml:"cloudWatchLoggingConfiguration"`
	// managed_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#managed_logging_configuration AwsAthenaWorkgroup#managed_logging_configuration}
	// Experimental.
	ManagedLoggingConfiguration *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty `field:"optional" json:"managedLoggingConfiguration" yaml:"managedLoggingConfiguration"`
	// s3_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#s3_logging_configuration AwsAthenaWorkgroup#s3_logging_configuration}
	// Experimental.
	S3LoggingConfiguration *AwsAthenaWorkgroup_S3LoggingConfigurationProperty `field:"optional" json:"s3LoggingConfiguration" yaml:"s3LoggingConfiguration"`
}

