package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationConfigurationProperty struct {
	// application_code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration TfApplication#application_code_configuration}
	// Experimental.
	ApplicationCodeConfiguration *TfApplication_ApplicationCodeConfigurationProperty `field:"required" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// application_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_encryption_configuration TfApplication#application_encryption_configuration}
	// Experimental.
	ApplicationEncryptionConfiguration *TfApplication_ApplicationEncryptionConfigurationProperty `field:"optional" json:"applicationEncryptionConfiguration" yaml:"applicationEncryptionConfiguration"`
	// application_snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration TfApplication#application_snapshot_configuration}
	// Experimental.
	ApplicationSnapshotConfiguration *TfApplication_ApplicationSnapshotConfigurationProperty `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// environment_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#environment_properties TfApplication#environment_properties}
	// Experimental.
	EnvironmentProperties *TfApplication_EnvironmentPropertiesProperty `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// flink_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration TfApplication#flink_application_configuration}
	// Experimental.
	FlinkApplicationConfiguration *TfApplication_FlinkApplicationConfigurationProperty `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#run_configuration TfApplication#run_configuration}
	// Experimental.
	RunConfiguration *TfApplication_RunConfigurationProperty `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// sql_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration TfApplication#sql_application_configuration}
	// Experimental.
	SqlApplicationConfiguration *TfApplication_SqlApplicationConfigurationProperty `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#vpc_configuration TfApplication#vpc_configuration}
	// Experimental.
	VpcConfiguration *TfApplication_VpcConfigurationProperty `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

