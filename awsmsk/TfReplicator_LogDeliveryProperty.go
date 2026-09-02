package awsmsk


// Experimental.
type TfReplicator_LogDeliveryProperty struct {
	// replicator_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#replicator_log_delivery TfReplicator#replicator_log_delivery}
	// Experimental.
	ReplicatorLogDelivery *TfReplicator_ReplicatorLogDeliveryProperty `field:"optional" json:"replicatorLogDelivery" yaml:"replicatorLogDelivery"`
}

