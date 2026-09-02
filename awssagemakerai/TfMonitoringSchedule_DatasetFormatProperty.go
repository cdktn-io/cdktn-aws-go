package awssagemakerai


// Experimental.
type TfMonitoringSchedule_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#csv TfMonitoringSchedule#csv}
	// Experimental.
	Csv *TfMonitoringSchedule_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#json TfMonitoringSchedule#json}
	// Experimental.
	Json *TfMonitoringSchedule_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

