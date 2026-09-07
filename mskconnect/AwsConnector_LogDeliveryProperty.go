package mskconnect


// Experimental.
type AwsConnector_LogDeliveryProperty struct {
	// worker_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#worker_log_delivery AwsConnector#worker_log_delivery}
	// Experimental.
	WorkerLogDelivery *AwsConnector_WorkerLogDeliveryProperty `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

