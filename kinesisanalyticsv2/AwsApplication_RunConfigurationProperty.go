package kinesisanalyticsv2


// Experimental.
type AwsApplication_RunConfigurationProperty struct {
	// application_restore_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_restore_configuration AwsApplication#application_restore_configuration}
	// Experimental.
	ApplicationRestoreConfiguration *AwsApplication_ApplicationRestoreConfigurationProperty `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// flink_run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_run_configuration AwsApplication#flink_run_configuration}
	// Experimental.
	FlinkRunConfiguration *AwsApplication_FlinkRunConfigurationProperty `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

