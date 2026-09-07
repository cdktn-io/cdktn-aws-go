package msk


// Experimental.
type AwsCluster_LoggingInfoProperty struct {
	// broker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#broker_logs AwsCluster#broker_logs}
	// Experimental.
	BrokerLogs *AwsCluster_BrokerLogsProperty `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

