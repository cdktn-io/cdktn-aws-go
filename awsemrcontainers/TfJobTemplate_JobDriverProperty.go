package awsemrcontainers


// Experimental.
type TfJobTemplate_JobDriverProperty struct {
	// spark_sql_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_sql_job_driver TfJobTemplate#spark_sql_job_driver}
	// Experimental.
	SparkSqlJobDriver *TfJobTemplate_SparkSqlJobDriverProperty `field:"optional" json:"sparkSqlJobDriver" yaml:"sparkSqlJobDriver"`
	// spark_submit_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_submit_job_driver TfJobTemplate#spark_submit_job_driver}
	// Experimental.
	SparkSubmitJobDriver *TfJobTemplate_SparkSubmitJobDriverProperty `field:"optional" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}

