package sagemakerai


// Experimental.
type AwsTrainingJob_ProfilerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#disable_profiler AwsTrainingJob#disable_profiler}.
	// Experimental.
	DisableProfiler interface{} `field:"optional" json:"disableProfiler" yaml:"disableProfiler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#profiling_interval_in_milliseconds AwsTrainingJob#profiling_interval_in_milliseconds}.
	// Experimental.
	ProfilingIntervalInMilliseconds *float64 `field:"optional" json:"profilingIntervalInMilliseconds" yaml:"profilingIntervalInMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#profiling_parameters AwsTrainingJob#profiling_parameters}.
	// Experimental.
	ProfilingParameters *map[string]*string `field:"optional" json:"profilingParameters" yaml:"profilingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#s3_output_path AwsTrainingJob#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

