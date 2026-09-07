package backup


// Experimental.
type AwsPlan_AdvancedBackupSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#backup_options AwsPlan#backup_options}.
	// Experimental.
	BackupOptions *map[string]*string `field:"required" json:"backupOptions" yaml:"backupOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#resource_type AwsPlan#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

