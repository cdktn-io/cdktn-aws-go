package emrcontainers


// Experimental.
type AwsJobTemplate_JobDriverProperty struct {
	// spark_sql_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_sql_job_driver AwsJobTemplate#spark_sql_job_driver}
	// Experimental.
	SparkSqlJobDriver *AwsJobTemplate_SparkSqlJobDriverProperty `field:"optional" json:"sparkSqlJobDriver" yaml:"sparkSqlJobDriver"`
	// spark_submit_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_submit_job_driver AwsJobTemplate#spark_submit_job_driver}
	// Experimental.
	SparkSubmitJobDriver *AwsJobTemplate_SparkSubmitJobDriverProperty `field:"optional" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}

