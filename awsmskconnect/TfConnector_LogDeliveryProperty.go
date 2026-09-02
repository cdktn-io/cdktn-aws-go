package awsmskconnect


// Experimental.
type TfConnector_LogDeliveryProperty struct {
	// worker_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#worker_log_delivery TfConnector#worker_log_delivery}
	// Experimental.
	WorkerLogDelivery *TfConnector_WorkerLogDeliveryProperty `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

