package awsbackup


// Experimental.
type TfPlan_ScanSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#malware_scanner TfPlan#malware_scanner}.
	// Experimental.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#resource_types TfPlan#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scanner_role_arn TfPlan#scanner_role_arn}.
	// Experimental.
	ScannerRoleArn *string `field:"required" json:"scannerRoleArn" yaml:"scannerRoleArn"`
}

