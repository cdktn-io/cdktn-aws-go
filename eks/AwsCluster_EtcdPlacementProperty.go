package eks


// Experimental.
type AwsCluster_EtcdPlacementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#spread_level AwsCluster#spread_level}.
	// Experimental.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

