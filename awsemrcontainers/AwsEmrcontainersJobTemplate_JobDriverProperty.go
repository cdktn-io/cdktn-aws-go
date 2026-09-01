package awsemrcontainers


// Experimental.
type AwsEmrcontainersJobTemplate_JobDriverProperty struct {
	// spark_sql_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_sql_job_driver AwsEmrcontainersJobTemplate#spark_sql_job_driver}
	// Experimental.
	SparkSqlJobDriver *AwsEmrcontainersJobTemplate_SparkSqlJobDriverProperty `field:"optional" json:"sparkSqlJobDriver" yaml:"sparkSqlJobDriver"`
	// spark_submit_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_submit_job_driver AwsEmrcontainersJobTemplate#spark_submit_job_driver}
	// Experimental.
	SparkSubmitJobDriver *AwsEmrcontainersJobTemplate_SparkSubmitJobDriverProperty `field:"optional" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}

