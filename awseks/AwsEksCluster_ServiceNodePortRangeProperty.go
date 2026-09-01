package awseks


// Experimental.
type AwsEksCluster_ServiceNodePortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#max_port AwsEksCluster#max_port}.
	// Experimental.
	MaxPort *float64 `field:"optional" json:"maxPort" yaml:"maxPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#min_port AwsEksCluster#min_port}.
	// Experimental.
	MinPort *float64 `field:"optional" json:"minPort" yaml:"minPort"`
}

