package mskconnect


// Experimental.
type AwsConnector_ProvisionedCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#worker_count AwsConnector#worker_count}.
	// Experimental.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#mcu_count AwsConnector#mcu_count}.
	// Experimental.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

