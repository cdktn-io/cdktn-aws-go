package awsmsk


// Experimental.
type AwsMskCluster_LoggingInfoProperty struct {
	// broker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#broker_logs AwsMskCluster#broker_logs}
	// Experimental.
	BrokerLogs *AwsMskCluster_BrokerLogsProperty `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

