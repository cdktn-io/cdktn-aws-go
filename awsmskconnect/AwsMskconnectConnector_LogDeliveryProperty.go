package awsmskconnect


// Experimental.
type AwsMskconnectConnector_LogDeliveryProperty struct {
	// worker_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#worker_log_delivery AwsMskconnectConnector#worker_log_delivery}
	// Experimental.
	WorkerLogDelivery *AwsMskconnectConnector_WorkerLogDeliveryProperty `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

