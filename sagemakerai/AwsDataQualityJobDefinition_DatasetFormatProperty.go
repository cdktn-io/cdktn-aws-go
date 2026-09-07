package sagemakerai


// Experimental.
type AwsDataQualityJobDefinition_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#csv AwsDataQualityJobDefinition#csv}
	// Experimental.
	Csv *AwsDataQualityJobDefinition_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#json AwsDataQualityJobDefinition#json}
	// Experimental.
	Json *AwsDataQualityJobDefinition_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

