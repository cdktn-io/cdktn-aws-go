package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty struct {
	// checkpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#checkpoint_configuration AwsKinesisanalyticsv2Application#checkpoint_configuration}
	// Experimental.
	CheckpointConfiguration *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration AwsKinesisanalyticsv2Application#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// parallelism_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#parallelism_configuration AwsKinesisanalyticsv2Application#parallelism_configuration}
	// Experimental.
	ParallelismConfiguration *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

