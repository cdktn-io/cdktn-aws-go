package awsopensearch


// Experimental.
type AwsOpensearchDomain_NodeOptionsProperty struct {
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_config AwsOpensearchDomain#node_config}
	// Experimental.
	NodeConfig *AwsOpensearchDomain_NodeConfigProperty `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_type AwsOpensearchDomain#node_type}.
	// Experimental.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

