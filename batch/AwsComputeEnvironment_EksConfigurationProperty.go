package batch


// Experimental.
type AwsComputeEnvironment_EksConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#eks_cluster_arn AwsComputeEnvironment#eks_cluster_arn}.
	// Experimental.
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#kubernetes_namespace AwsComputeEnvironment#kubernetes_namespace}.
	// Experimental.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

