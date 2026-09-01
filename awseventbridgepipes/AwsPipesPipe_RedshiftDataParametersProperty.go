package awseventbridgepipes


// Experimental.
type AwsPipesPipe_RedshiftDataParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#database AwsPipesPipe#database}.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqls AwsPipesPipe#sqls}.
	// Experimental.
	Sqls *[]*string `field:"required" json:"sqls" yaml:"sqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#db_user AwsPipesPipe#db_user}.
	// Experimental.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#secret_manager_arn AwsPipesPipe#secret_manager_arn}.
	// Experimental.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#statement_name AwsPipesPipe#statement_name}.
	// Experimental.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#with_event AwsPipesPipe#with_event}.
	// Experimental.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

