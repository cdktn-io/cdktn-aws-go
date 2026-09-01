package awsmsk


// Experimental.
type AwsMskCluster_OpenMonitoringProperty struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#prometheus AwsMskCluster#prometheus}
	// Experimental.
	Prometheus *AwsMskCluster_PrometheusProperty `field:"required" json:"prometheus" yaml:"prometheus"`
}

