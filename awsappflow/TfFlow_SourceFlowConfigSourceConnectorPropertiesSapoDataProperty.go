package awsappflow


// Experimental.
type TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object_path TfFlow#object_path}.
	// Experimental.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// pagination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#pagination_config TfFlow#pagination_config}
	// Experimental.
	PaginationConfig *TfFlow_PaginationConfigProperty `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// parallelism_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#parallelism_config TfFlow#parallelism_config}
	// Experimental.
	ParallelismConfig *TfFlow_ParallelismConfigProperty `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

