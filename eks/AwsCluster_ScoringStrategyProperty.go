package eks


// Experimental.
type AwsCluster_ScoringStrategyProperty struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resource AwsCluster#resource}
	// Experimental.
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#type AwsCluster#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

