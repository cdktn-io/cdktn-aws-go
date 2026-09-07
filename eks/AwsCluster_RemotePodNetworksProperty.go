package eks


// Experimental.
type AwsCluster_RemotePodNetworksProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#cidrs AwsCluster#cidrs}.
	// Experimental.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

