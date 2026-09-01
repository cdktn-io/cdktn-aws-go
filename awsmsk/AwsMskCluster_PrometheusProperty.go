package awsmsk


// Experimental.
type AwsMskCluster_PrometheusProperty struct {
	// jmx_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#jmx_exporter AwsMskCluster#jmx_exporter}
	// Experimental.
	JmxExporter *AwsMskCluster_JmxExporterProperty `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// node_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#node_exporter AwsMskCluster#node_exporter}
	// Experimental.
	NodeExporter *AwsMskCluster_NodeExporterProperty `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

