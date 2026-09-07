package msk


// Experimental.
type AwsCluster_PrometheusProperty struct {
	// jmx_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#jmx_exporter AwsCluster#jmx_exporter}
	// Experimental.
	JmxExporter *AwsCluster_JmxExporterProperty `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// node_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#node_exporter AwsCluster#node_exporter}
	// Experimental.
	NodeExporter *AwsCluster_NodeExporterProperty `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

