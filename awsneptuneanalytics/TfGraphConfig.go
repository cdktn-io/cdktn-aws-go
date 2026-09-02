package awsneptuneanalytics

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGraphConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#provisioned_memory TfGraph#provisioned_memory}
	// Experimental.
	ProvisionedMemory *float64 `field:"required" json:"provisionedMemory" yaml:"provisionedMemory"`
	// A value that indicates whether the graph has deletion protection enabled.
	//
	// The graph can't be deleted when deletion protection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#deletion_protection TfGraph#deletion_protection}
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The graph name.
	//
	// For example: my-graph-1.
	// 								The name must contain from 1 to 63 letters, numbers, or hyphens,
	// 								and its first character must be a letter. It cannot end with a hyphen or contain two consecutive hyphens.
	// 								If you don't specify a graph name, a unique graph name is generated for you using the prefix graph-for,
	// 								followed by a combination of Stack Name and a UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#graph_name TfGraph#graph_name}
	// Experimental.
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// Allows user to specify name prefix and have remainder of name automatically generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#graph_name_prefix TfGraph#graph_name_prefix}
	// Experimental.
	GraphNamePrefix *string `field:"optional" json:"graphNamePrefix" yaml:"graphNamePrefix"`
	// Specifies a KMS key to use to encrypt data in the new graph.
	//
	// Value must be ARN of KMS Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#kms_key_identifier TfGraph#kms_key_identifier}
	// Experimental.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Specifies whether or not the graph can be reachable over the internet.
	//
	// All access to graphs is IAM authenticated.
	// 								When the graph is publicly available, its domain name system (DNS) endpoint resolves to
	// 								the public IP address from the internet. When the graph isn't publicly available, you need
	// 								to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private
	// 								IP address that is reachable from the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#public_connectivity TfGraph#public_connectivity}
	// Experimental.
	PublicConnectivity interface{} `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#region TfGraph#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The number of replicas in other AZs.  Value must be between 0 and 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#replica_count TfGraph#replica_count}
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#tags TfGraph#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#timeouts TfGraph#timeouts}
	// Experimental.
	Timeouts *TfGraph_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vector_search_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptunegraph_graph#vector_search_configuration TfGraph#vector_search_configuration}
	// Experimental.
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

