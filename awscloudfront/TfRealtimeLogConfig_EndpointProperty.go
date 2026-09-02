package awscloudfront


// Experimental.
type TfRealtimeLogConfig_EndpointProperty struct {
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_realtime_log_config#kinesis_stream_config TfRealtimeLogConfig#kinesis_stream_config}
	// Experimental.
	KinesisStreamConfig *TfRealtimeLogConfig_KinesisStreamConfigProperty `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_realtime_log_config#stream_type TfRealtimeLogConfig#stream_type}.
	// Experimental.
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}

