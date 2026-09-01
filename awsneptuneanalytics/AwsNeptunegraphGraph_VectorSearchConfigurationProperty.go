package awsneptuneanalytics


// Experimental.
type AwsNeptunegraphGraph_VectorSearchConfigurationProperty struct {
	// Specifies the number of dimensions for vector embeddings.  Value must be between 1 and 65,535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#vector_search_dimension AwsNeptunegraphGraph#vector_search_dimension}
	// Experimental.
	VectorSearchDimension *float64 `field:"optional" json:"vectorSearchDimension" yaml:"vectorSearchDimension"`
}

