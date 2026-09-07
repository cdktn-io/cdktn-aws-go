package fsx


// Experimental.
type AwsOntapVolume_SnaplockConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#snaplock_type AwsOntapVolume#snaplock_type}.
	// Experimental.
	SnaplockType *string `field:"required" json:"snaplockType" yaml:"snaplockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#audit_log_volume AwsOntapVolume#audit_log_volume}.
	// Experimental.
	AuditLogVolume interface{} `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// autocommit_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#autocommit_period AwsOntapVolume#autocommit_period}
	// Experimental.
	AutocommitPeriod *AwsOntapVolume_AutocommitPeriodProperty `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#privileged_delete AwsOntapVolume#privileged_delete}.
	// Experimental.
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// retention_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#retention_period AwsOntapVolume#retention_period}
	// Experimental.
	RetentionPeriod *AwsOntapVolume_RetentionPeriodProperty `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#volume_append_mode_enabled AwsOntapVolume#volume_append_mode_enabled}.
	// Experimental.
	VolumeAppendModeEnabled interface{} `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

