package awsemrcontainers


// Experimental.
type AwsEmrcontainersJobTemplate_SparkSqlJobDriverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#entry_point AwsEmrcontainersJobTemplate#entry_point}.
	// Experimental.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_sql_parameters AwsEmrcontainersJobTemplate#spark_sql_parameters}.
	// Experimental.
	SparkSqlParameters *string `field:"optional" json:"sparkSqlParameters" yaml:"sparkSqlParameters"`
}

