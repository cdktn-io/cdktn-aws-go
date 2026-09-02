package awseks


// Experimental.
type TfCluster_RemoteNodeNetworksProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#cidrs TfCluster#cidrs}.
	// Experimental.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

