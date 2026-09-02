package awssagemakerai


// Experimental.
type TfDataQualityJobDefinition_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#csv TfDataQualityJobDefinition#csv}
	// Experimental.
	Csv *TfDataQualityJobDefinition_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#json TfDataQualityJobDefinition#json}
	// Experimental.
	Json *TfDataQualityJobDefinition_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

