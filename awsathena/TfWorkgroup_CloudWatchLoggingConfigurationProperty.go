package awsathena


// Experimental.
type TfWorkgroup_CloudWatchLoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enabled TfWorkgroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#log_group TfWorkgroup#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#log_stream_name_prefix TfWorkgroup#log_stream_name_prefix}.
	// Experimental.
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// log_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#log_type TfWorkgroup#log_type}
	// Experimental.
	LogType interface{} `field:"optional" json:"logType" yaml:"logType"`
}

