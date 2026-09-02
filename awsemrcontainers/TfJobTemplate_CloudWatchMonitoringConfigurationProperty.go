package awsemrcontainers


// Experimental.
type TfJobTemplate_CloudWatchMonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#log_group_name TfJobTemplate#log_group_name}.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#log_stream_name_prefix TfJobTemplate#log_stream_name_prefix}.
	// Experimental.
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
}

