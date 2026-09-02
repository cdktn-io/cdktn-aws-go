package awseventbridge


// Experimental.
type TfTarget_RedshiftTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#database TfTarget#database}.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#db_user TfTarget#db_user}.
	// Experimental.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#secrets_manager_arn TfTarget#secrets_manager_arn}.
	// Experimental.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#sql TfTarget#sql}.
	// Experimental.
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#statement_name TfTarget#statement_name}.
	// Experimental.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#with_event TfTarget#with_event}.
	// Experimental.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

