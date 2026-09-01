package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#csv AwsSagemakerDataQualityJobDefinition#csv}
	// Experimental.
	Csv *AwsSagemakerDataQualityJobDefinition_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#json AwsSagemakerDataQualityJobDefinition#json}
	// Experimental.
	Json *AwsSagemakerDataQualityJobDefinition_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

