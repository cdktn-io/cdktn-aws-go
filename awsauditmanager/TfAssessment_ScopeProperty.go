package awsauditmanager


// Experimental.
type TfAssessment_ScopeProperty struct {
	// aws_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#aws_accounts TfAssessment#aws_accounts}
	// Experimental.
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// aws_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#aws_services TfAssessment#aws_services}
	// Experimental.
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

