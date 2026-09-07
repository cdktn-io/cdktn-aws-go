package opensearch


// Experimental.
type AwsDomain_NodeOptionsProperty struct {
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_config AwsDomain#node_config}
	// Experimental.
	NodeConfig *AwsDomain_NodeConfigProperty `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_type AwsDomain#node_type}.
	// Experimental.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

