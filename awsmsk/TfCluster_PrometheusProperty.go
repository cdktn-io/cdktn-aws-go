package awsmsk


// Experimental.
type TfCluster_PrometheusProperty struct {
	// jmx_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#jmx_exporter TfCluster#jmx_exporter}
	// Experimental.
	JmxExporter *TfCluster_JmxExporterProperty `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// node_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#node_exporter TfCluster#node_exporter}
	// Experimental.
	NodeExporter *TfCluster_NodeExporterProperty `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

