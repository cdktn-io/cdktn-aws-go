package emrcontainers


// Experimental.
type AwsJobTemplate_SparkSubmitJobDriverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#entry_point AwsJobTemplate#entry_point}.
	// Experimental.
	EntryPoint *string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#entry_point_arguments AwsJobTemplate#entry_point_arguments}.
	// Experimental.
	EntryPointArguments *[]*string `field:"optional" json:"entryPointArguments" yaml:"entryPointArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#spark_submit_parameters AwsJobTemplate#spark_submit_parameters}.
	// Experimental.
	SparkSubmitParameters *string `field:"optional" json:"sparkSubmitParameters" yaml:"sparkSubmitParameters"`
}

