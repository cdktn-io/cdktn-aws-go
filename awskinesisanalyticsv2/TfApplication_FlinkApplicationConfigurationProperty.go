package awskinesisanalyticsv2


// Experimental.
type TfApplication_FlinkApplicationConfigurationProperty struct {
	// checkpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#checkpoint_configuration TfApplication#checkpoint_configuration}
	// Experimental.
	CheckpointConfiguration *TfApplication_CheckpointConfigurationProperty `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration TfApplication#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *TfApplication_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// parallelism_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#parallelism_configuration TfApplication#parallelism_configuration}
	// Experimental.
	ParallelismConfiguration *TfApplication_ParallelismConfigurationProperty `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

