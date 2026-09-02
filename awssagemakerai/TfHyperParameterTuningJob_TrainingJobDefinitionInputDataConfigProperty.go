package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#channel_name TfHyperParameterTuningJob#channel_name}.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#compression_type TfHyperParameterTuningJob#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#content_type TfHyperParameterTuningJob#content_type}.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#data_source TfHyperParameterTuningJob#data_source}
	// Experimental.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#input_mode TfHyperParameterTuningJob#input_mode}.
	// Experimental.
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#record_wrapper_type TfHyperParameterTuningJob#record_wrapper_type}.
	// Experimental.
	RecordWrapperType *string `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// shuffle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#shuffle_config TfHyperParameterTuningJob#shuffle_config}
	// Experimental.
	ShuffleConfig interface{} `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

