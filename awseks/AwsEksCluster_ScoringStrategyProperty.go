package awseks


// Experimental.
type AwsEksCluster_ScoringStrategyProperty struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resource AwsEksCluster#resource}
	// Experimental.
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#type AwsEksCluster#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

