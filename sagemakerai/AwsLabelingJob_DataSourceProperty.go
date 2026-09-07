package sagemakerai


// Experimental.
type AwsLabelingJob_DataSourceProperty struct {
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#s3_data_source AwsLabelingJob#s3_data_source}
	// Experimental.
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
	// sns_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#sns_data_source AwsLabelingJob#sns_data_source}
	// Experimental.
	SnsDataSource interface{} `field:"optional" json:"snsDataSource" yaml:"snsDataSource"`
}

