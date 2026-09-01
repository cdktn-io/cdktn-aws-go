package awsdrs


// Experimental.
type AwsDrsReplicationConfigurationTemplate_PitPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#interval AwsDrsReplicationConfigurationTemplate#interval}.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#retention_duration AwsDrsReplicationConfigurationTemplate#retention_duration}.
	// Experimental.
	RetentionDuration *float64 `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#units AwsDrsReplicationConfigurationTemplate#units}.
	// Experimental.
	Units *string `field:"required" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#enabled AwsDrsReplicationConfigurationTemplate#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#rule_id AwsDrsReplicationConfigurationTemplate#rule_id}.
	// Experimental.
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
}

