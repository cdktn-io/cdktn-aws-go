package awsmsk


// Experimental.
type TfCluster_OpenMonitoringProperty struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#prometheus TfCluster#prometheus}
	// Experimental.
	Prometheus *TfCluster_PrometheusProperty `field:"required" json:"prometheus" yaml:"prometheus"`
}

