package awschimesdkmediapipelines


// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_ElementsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// amazon_transcribe_call_analytics_processor_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#amazon_transcribe_call_analytics_processor_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#amazon_transcribe_call_analytics_processor_configuration}
	// Experimental.
	AmazonTranscribeCallAnalyticsProcessorConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty `field:"optional" json:"amazonTranscribeCallAnalyticsProcessorConfiguration" yaml:"amazonTranscribeCallAnalyticsProcessorConfiguration"`
	// amazon_transcribe_processor_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#amazon_transcribe_processor_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#amazon_transcribe_processor_configuration}
	// Experimental.
	AmazonTranscribeProcessorConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty `field:"optional" json:"amazonTranscribeProcessorConfiguration" yaml:"amazonTranscribeProcessorConfiguration"`
	// kinesis_data_stream_sink_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#kinesis_data_stream_sink_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#kinesis_data_stream_sink_configuration}
	// Experimental.
	KinesisDataStreamSinkConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty `field:"optional" json:"kinesisDataStreamSinkConfiguration" yaml:"kinesisDataStreamSinkConfiguration"`
	// lambda_function_sink_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#lambda_function_sink_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#lambda_function_sink_configuration}
	// Experimental.
	LambdaFunctionSinkConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty `field:"optional" json:"lambdaFunctionSinkConfiguration" yaml:"lambdaFunctionSinkConfiguration"`
	// s3_recording_sink_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#s3_recording_sink_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#s3_recording_sink_configuration}
	// Experimental.
	S3RecordingSinkConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty `field:"optional" json:"s3RecordingSinkConfiguration" yaml:"s3RecordingSinkConfiguration"`
	// sns_topic_sink_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sns_topic_sink_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#sns_topic_sink_configuration}
	// Experimental.
	SnsTopicSinkConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty `field:"optional" json:"snsTopicSinkConfiguration" yaml:"snsTopicSinkConfiguration"`
	// sqs_queue_sink_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sqs_queue_sink_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#sqs_queue_sink_configuration}
	// Experimental.
	SqsQueueSinkConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty `field:"optional" json:"sqsQueueSinkConfiguration" yaml:"sqsQueueSinkConfiguration"`
	// voice_analytics_processor_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#voice_analytics_processor_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#voice_analytics_processor_configuration}
	// Experimental.
	VoiceAnalyticsProcessorConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty `field:"optional" json:"voiceAnalyticsProcessorConfiguration" yaml:"voiceAnalyticsProcessorConfiguration"`
}

