package sagemakerai


// Experimental.
type AwsMonitoringSchedule_DatasetFormatProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#csv AwsMonitoringSchedule#csv}
	// Experimental.
	Csv *AwsMonitoringSchedule_CsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#json AwsMonitoringSchedule#json}
	// Experimental.
	Json *AwsMonitoringSchedule_JsonProperty `field:"optional" json:"json" yaml:"json"`
}

