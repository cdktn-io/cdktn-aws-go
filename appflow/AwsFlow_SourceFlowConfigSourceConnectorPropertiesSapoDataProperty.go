package appflow


// Experimental.
type AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object_path AwsFlow#object_path}.
	// Experimental.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// pagination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#pagination_config AwsFlow#pagination_config}
	// Experimental.
	PaginationConfig *AwsFlow_PaginationConfigProperty `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// parallelism_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#parallelism_config AwsFlow#parallelism_config}
	// Experimental.
	ParallelismConfig *AwsFlow_ParallelismConfigProperty `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

