package awsappflow


// Experimental.
type AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object_path AwsAppflowFlow#object_path}.
	// Experimental.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// pagination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#pagination_config AwsAppflowFlow#pagination_config}
	// Experimental.
	PaginationConfig *AwsAppflowFlow_PaginationConfigProperty `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// parallelism_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#parallelism_config AwsAppflowFlow#parallelism_config}
	// Experimental.
	ParallelismConfig *AwsAppflowFlow_ParallelismConfigProperty `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

