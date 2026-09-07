package kinesisanalyticsv2


// Experimental.
type AwsApplication_FlinkApplicationConfigurationProperty struct {
	// checkpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#checkpoint_configuration AwsApplication#checkpoint_configuration}
	// Experimental.
	CheckpointConfiguration *AwsApplication_CheckpointConfigurationProperty `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration AwsApplication#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *AwsApplication_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// parallelism_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#parallelism_configuration AwsApplication#parallelism_configuration}
	// Experimental.
	ParallelismConfiguration *AwsApplication_ParallelismConfigurationProperty `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

