package drs


// Experimental.
type AwsReplicationConfigurationTemplate_PitPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#interval AwsReplicationConfigurationTemplate#interval}.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#retention_duration AwsReplicationConfigurationTemplate#retention_duration}.
	// Experimental.
	RetentionDuration *float64 `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#units AwsReplicationConfigurationTemplate#units}.
	// Experimental.
	Units *string `field:"required" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#enabled AwsReplicationConfigurationTemplate#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#rule_id AwsReplicationConfigurationTemplate#rule_id}.
	// Experimental.
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
}

