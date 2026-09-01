package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_RunConfigurationProperty struct {
	// application_restore_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_restore_configuration AwsKinesisanalyticsv2Application#application_restore_configuration}
	// Experimental.
	ApplicationRestoreConfiguration *AwsKinesisanalyticsv2Application_ApplicationRestoreConfigurationProperty `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// flink_run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_run_configuration AwsKinesisanalyticsv2Application#flink_run_configuration}
	// Experimental.
	FlinkRunConfiguration *AwsKinesisanalyticsv2Application_FlinkRunConfigurationProperty `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

