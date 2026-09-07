package emrcontainers


// Experimental.
type AwsJobTemplate_SparkSqlJobDriverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#entry_point AwsJobTemplate#entry_point}.
	// Experimental.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_sql_parameters AwsJobTemplate#spark_sql_parameters}.
	// Experimental.
	SparkSqlParameters *string `field:"optional" json:"sparkSqlParameters" yaml:"sparkSqlParameters"`
}

