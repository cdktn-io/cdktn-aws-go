package awsopensearch


// Experimental.
type TfDomain_NodeOptionsProperty struct {
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_config TfDomain#node_config}
	// Experimental.
	NodeConfig *TfDomain_NodeConfigProperty `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_type TfDomain#node_type}.
	// Experimental.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

