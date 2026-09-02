package awsmsk


// Experimental.
type TfCluster_LoggingInfoProperty struct {
	// broker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#broker_logs TfCluster#broker_logs}
	// Experimental.
	BrokerLogs *TfCluster_BrokerLogsProperty `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

