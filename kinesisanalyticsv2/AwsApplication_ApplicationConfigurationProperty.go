package kinesisanalyticsv2


// Experimental.
type AwsApplication_ApplicationConfigurationProperty struct {
	// application_code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration AwsApplication#application_code_configuration}
	// Experimental.
	ApplicationCodeConfiguration *AwsApplication_ApplicationCodeConfigurationProperty `field:"required" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// application_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_encryption_configuration AwsApplication#application_encryption_configuration}
	// Experimental.
	ApplicationEncryptionConfiguration *AwsApplication_ApplicationEncryptionConfigurationProperty `field:"optional" json:"applicationEncryptionConfiguration" yaml:"applicationEncryptionConfiguration"`
	// application_snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration AwsApplication#application_snapshot_configuration}
	// Experimental.
	ApplicationSnapshotConfiguration *AwsApplication_ApplicationSnapshotConfigurationProperty `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// environment_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#environment_properties AwsApplication#environment_properties}
	// Experimental.
	EnvironmentProperties *AwsApplication_EnvironmentPropertiesProperty `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// flink_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration AwsApplication#flink_application_configuration}
	// Experimental.
	FlinkApplicationConfiguration *AwsApplication_FlinkApplicationConfigurationProperty `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#run_configuration AwsApplication#run_configuration}
	// Experimental.
	RunConfiguration *AwsApplication_RunConfigurationProperty `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// sql_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration AwsApplication#sql_application_configuration}
	// Experimental.
	SqlApplicationConfiguration *AwsApplication_SqlApplicationConfigurationProperty `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#vpc_configuration AwsApplication#vpc_configuration}
	// Experimental.
	VpcConfiguration *AwsApplication_VpcConfigurationProperty `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

