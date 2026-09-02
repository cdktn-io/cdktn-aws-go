package awskinesisanalyticsv2


// Experimental.
type TfApplication_RunConfigurationProperty struct {
	// application_restore_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_restore_configuration TfApplication#application_restore_configuration}
	// Experimental.
	ApplicationRestoreConfiguration *TfApplication_ApplicationRestoreConfigurationProperty `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// flink_run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_run_configuration TfApplication#flink_run_configuration}
	// Experimental.
	FlinkRunConfiguration *TfApplication_FlinkRunConfigurationProperty `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

