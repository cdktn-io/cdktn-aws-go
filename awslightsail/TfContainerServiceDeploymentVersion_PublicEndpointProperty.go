package awslightsail


// Experimental.
type TfContainerServiceDeploymentVersion_PublicEndpointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#container_name TfContainerServiceDeploymentVersion#container_name}.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#container_port TfContainerServiceDeploymentVersion#container_port}.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#health_check TfContainerServiceDeploymentVersion#health_check}
	// Experimental.
	HealthCheck *TfContainerServiceDeploymentVersion_HealthCheckProperty `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

