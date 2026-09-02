package awscloudfront


// Experimental.
type TfRealtimeLogConfig_KinesisStreamConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_realtime_log_config#role_arn TfRealtimeLogConfig#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_realtime_log_config#stream_arn TfRealtimeLogConfig#stream_arn}.
	// Experimental.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

