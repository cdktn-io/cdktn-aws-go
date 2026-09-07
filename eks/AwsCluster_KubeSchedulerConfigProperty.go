package eks


// Experimental.
type AwsCluster_KubeSchedulerConfigProperty struct {
	// node_resources_fit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#node_resources_fit AwsCluster#node_resources_fit}
	// Experimental.
	NodeResourcesFit *AwsCluster_NodeResourcesFitProperty `field:"optional" json:"nodeResourcesFit" yaml:"nodeResourcesFit"`
}

