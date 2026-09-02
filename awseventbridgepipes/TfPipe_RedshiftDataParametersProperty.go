package awseventbridgepipes


// Experimental.
type TfPipe_RedshiftDataParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#database TfPipe#database}.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqls TfPipe#sqls}.
	// Experimental.
	Sqls *[]*string `field:"required" json:"sqls" yaml:"sqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#db_user TfPipe#db_user}.
	// Experimental.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#secret_manager_arn TfPipe#secret_manager_arn}.
	// Experimental.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#statement_name TfPipe#statement_name}.
	// Experimental.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#with_event TfPipe#with_event}.
	// Experimental.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

