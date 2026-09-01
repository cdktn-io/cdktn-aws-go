package awseks


// Experimental.
type AwsEksCluster_KubernetesNetworkConfigProperty struct {
	// elastic_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#elastic_load_balancing AwsEksCluster#elastic_load_balancing}
	// Experimental.
	ElasticLoadBalancing *AwsEksCluster_ElasticLoadBalancingProperty `field:"optional" json:"elasticLoadBalancing" yaml:"elasticLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#ip_family AwsEksCluster#ip_family}.
	// Experimental.
	IpFamily *string `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#service_ipv4_cidr AwsEksCluster#service_ipv4_cidr}.
	// Experimental.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

