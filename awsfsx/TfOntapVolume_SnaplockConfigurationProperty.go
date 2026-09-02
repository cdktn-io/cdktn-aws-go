package awsfsx


// Experimental.
type TfOntapVolume_SnaplockConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#snaplock_type TfOntapVolume#snaplock_type}.
	// Experimental.
	SnaplockType *string `field:"required" json:"snaplockType" yaml:"snaplockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#audit_log_volume TfOntapVolume#audit_log_volume}.
	// Experimental.
	AuditLogVolume interface{} `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// autocommit_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#autocommit_period TfOntapVolume#autocommit_period}
	// Experimental.
	AutocommitPeriod *TfOntapVolume_AutocommitPeriodProperty `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#privileged_delete TfOntapVolume#privileged_delete}.
	// Experimental.
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// retention_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#retention_period TfOntapVolume#retention_period}
	// Experimental.
	RetentionPeriod *TfOntapVolume_RetentionPeriodProperty `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#volume_append_mode_enabled TfOntapVolume#volume_append_mode_enabled}.
	// Experimental.
	VolumeAppendModeEnabled interface{} `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

