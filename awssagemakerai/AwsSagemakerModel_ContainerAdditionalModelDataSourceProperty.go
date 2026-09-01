package awssagemakerai


// Experimental.
type AwsSagemakerModel_ContainerAdditionalModelDataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#channel_name AwsSagemakerModel#channel_name}.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#s3_data_source AwsSagemakerModel#s3_data_source}
	// Experimental.
	S3DataSource interface{} `field:"required" json:"s3DataSource" yaml:"s3DataSource"`
}

