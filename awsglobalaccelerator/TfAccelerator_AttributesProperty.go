package awsglobalaccelerator


// Experimental.
type TfAccelerator_AttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_accelerator#flow_logs_enabled TfAccelerator#flow_logs_enabled}.
	// Experimental.
	FlowLogsEnabled interface{} `field:"optional" json:"flowLogsEnabled" yaml:"flowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_accelerator#flow_logs_s3_bucket TfAccelerator#flow_logs_s3_bucket}.
	// Experimental.
	FlowLogsS3Bucket *string `field:"optional" json:"flowLogsS3Bucket" yaml:"flowLogsS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_accelerator#flow_logs_s3_prefix TfAccelerator#flow_logs_s3_prefix}.
	// Experimental.
	FlowLogsS3Prefix *string `field:"optional" json:"flowLogsS3Prefix" yaml:"flowLogsS3Prefix"`
}

