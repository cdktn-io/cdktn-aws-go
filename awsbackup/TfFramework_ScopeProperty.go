package awsbackup


// Experimental.
type TfFramework_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#compliance_resource_ids TfFramework#compliance_resource_ids}.
	// Experimental.
	ComplianceResourceIds *[]*string `field:"optional" json:"complianceResourceIds" yaml:"complianceResourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#compliance_resource_types TfFramework#compliance_resource_types}.
	// Experimental.
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#tags TfFramework#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

