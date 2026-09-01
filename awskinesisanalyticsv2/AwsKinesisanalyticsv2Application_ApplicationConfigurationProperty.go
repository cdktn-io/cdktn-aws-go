package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty struct {
	// application_code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration AwsKinesisanalyticsv2Application#application_code_configuration}
	// Experimental.
	ApplicationCodeConfiguration *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty `field:"required" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// application_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_encryption_configuration AwsKinesisanalyticsv2Application#application_encryption_configuration}
	// Experimental.
	ApplicationEncryptionConfiguration *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty `field:"optional" json:"applicationEncryptionConfiguration" yaml:"applicationEncryptionConfiguration"`
	// application_snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration AwsKinesisanalyticsv2Application#application_snapshot_configuration}
	// Experimental.
	ApplicationSnapshotConfiguration *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// environment_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#environment_properties AwsKinesisanalyticsv2Application#environment_properties}
	// Experimental.
	EnvironmentProperties *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// flink_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration AwsKinesisanalyticsv2Application#flink_application_configuration}
	// Experimental.
	FlinkApplicationConfiguration *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#run_configuration AwsKinesisanalyticsv2Application#run_configuration}
	// Experimental.
	RunConfiguration *AwsKinesisanalyticsv2Application_RunConfigurationProperty `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// sql_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration AwsKinesisanalyticsv2Application#sql_application_configuration}
	// Experimental.
	SqlApplicationConfiguration *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#vpc_configuration AwsKinesisanalyticsv2Application#vpc_configuration}
	// Experimental.
	VpcConfiguration *AwsKinesisanalyticsv2Application_VpcConfigurationProperty `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

