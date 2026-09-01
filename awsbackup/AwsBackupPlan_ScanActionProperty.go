package awsbackup


// Experimental.
type AwsBackupPlan_ScanActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#malware_scanner AwsBackupPlan#malware_scanner}.
	// Experimental.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scan_mode AwsBackupPlan#scan_mode}.
	// Experimental.
	ScanMode *string `field:"required" json:"scanMode" yaml:"scanMode"`
}

