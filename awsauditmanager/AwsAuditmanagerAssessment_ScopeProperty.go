package awsauditmanager


// Experimental.
type AwsAuditmanagerAssessment_ScopeProperty struct {
	// aws_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#aws_accounts AwsAuditmanagerAssessment#aws_accounts}
	// Experimental.
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// aws_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#aws_services AwsAuditmanagerAssessment#aws_services}
	// Experimental.
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

