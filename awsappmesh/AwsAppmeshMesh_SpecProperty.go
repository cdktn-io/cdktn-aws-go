package awsappmesh


// Experimental.
type AwsAppmeshMesh_SpecProperty struct {
	// egress_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_mesh#egress_filter AwsAppmeshMesh#egress_filter}
	// Experimental.
	EgressFilter *AwsAppmeshMesh_EgressFilterProperty `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_mesh#service_discovery AwsAppmeshMesh#service_discovery}
	// Experimental.
	ServiceDiscovery *AwsAppmeshMesh_ServiceDiscoveryProperty `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

