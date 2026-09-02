package awscodedeploy


// Experimental.
type TfDeploymentGroup_EcsServiceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#cluster_name TfDeploymentGroup#cluster_name}.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#service_name TfDeploymentGroup#service_name}.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

