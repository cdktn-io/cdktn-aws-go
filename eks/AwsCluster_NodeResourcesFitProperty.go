package eks


// Experimental.
type AwsCluster_NodeResourcesFitProperty struct {
	// scoring_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#scoring_strategy AwsCluster#scoring_strategy}
	// Experimental.
	ScoringStrategy *AwsCluster_ScoringStrategyProperty `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
}

