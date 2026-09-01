package awslightsail


// Experimental.
type AwsLightsailContainerServiceDeploymentVersion_PublicEndpointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#container_name AwsLightsailContainerServiceDeploymentVersion#container_name}.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#container_port AwsLightsailContainerServiceDeploymentVersion#container_port}.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#health_check AwsLightsailContainerServiceDeploymentVersion#health_check}
	// Experimental.
	HealthCheck *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

