package awskinesisfirehose


// Experimental.
type TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProcessorsParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#parameter_name TfDeliveryStream#parameter_name}.
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#parameter_value TfDeliveryStream#parameter_value}.
	// Experimental.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

