package eks


// Experimental.
type AwsCluster_ComputeConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#enabled AwsCluster#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#node_pools AwsCluster#node_pools}.
	// Experimental.
	NodePools *[]*string `field:"optional" json:"nodePools" yaml:"nodePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#node_role_arn AwsCluster#node_role_arn}.
	// Experimental.
	NodeRoleArn *string `field:"optional" json:"nodeRoleArn" yaml:"nodeRoleArn"`
}

