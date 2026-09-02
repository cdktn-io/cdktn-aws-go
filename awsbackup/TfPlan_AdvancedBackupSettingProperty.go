package awsbackup


// Experimental.
type TfPlan_AdvancedBackupSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#backup_options TfPlan#backup_options}.
	// Experimental.
	BackupOptions *map[string]*string `field:"required" json:"backupOptions" yaml:"backupOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#resource_type TfPlan#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

