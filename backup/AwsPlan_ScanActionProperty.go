package backup


// Experimental.
type AwsPlan_ScanActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#malware_scanner AwsPlan#malware_scanner}.
	// Experimental.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#scan_mode AwsPlan#scan_mode}.
	// Experimental.
	ScanMode *string `field:"required" json:"scanMode" yaml:"scanMode"`
}

