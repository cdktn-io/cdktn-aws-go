package backup


// Experimental.
type AwsPlan_ScanSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#malware_scanner AwsPlan#malware_scanner}.
	// Experimental.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#resource_types AwsPlan#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scanner_role_arn AwsPlan#scanner_role_arn}.
	// Experimental.
	ScannerRoleArn *string `field:"required" json:"scannerRoleArn" yaml:"scannerRoleArn"`
}

