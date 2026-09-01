package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#csv AwsSagemakerMonitoringSchedule#csv}
	// Experimental.
	Csv *AwsSagemakerMonitoringSchedule_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#json AwsSagemakerMonitoringSchedule#json}
	// Experimental.
	Json *AwsSagemakerMonitoringSchedule_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

