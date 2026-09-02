package awsappmesh


// Experimental.
type TfMesh_SpecProperty struct {
	// egress_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_mesh#egress_filter TfMesh#egress_filter}
	// Experimental.
	EgressFilter *TfMesh_EgressFilterProperty `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_mesh#service_discovery TfMesh#service_discovery}
	// Experimental.
	ServiceDiscovery *TfMesh_ServiceDiscoveryProperty `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

