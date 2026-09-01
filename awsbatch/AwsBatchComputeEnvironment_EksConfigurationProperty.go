package awsbatch


// Experimental.
type AwsBatchComputeEnvironment_EksConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#eks_cluster_arn AwsBatchComputeEnvironment#eks_cluster_arn}.
	// Experimental.
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#kubernetes_namespace AwsBatchComputeEnvironment#kubernetes_namespace}.
	// Experimental.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

