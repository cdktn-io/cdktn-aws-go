package awseks


// Experimental.
type TfCluster_NodeResourcesFitProperty struct {
	// scoring_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#scoring_strategy TfCluster#scoring_strategy}
	// Experimental.
	ScoringStrategy *TfCluster_ScoringStrategyProperty `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
}

