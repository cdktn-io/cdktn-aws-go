package lightsail


// Experimental.
type AwsContainerServiceDeploymentVersion_HealthCheckProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#healthy_threshold AwsContainerServiceDeploymentVersion#healthy_threshold}.
	// Experimental.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#interval_seconds AwsContainerServiceDeploymentVersion#interval_seconds}.
	// Experimental.
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#path AwsContainerServiceDeploymentVersion#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#success_codes AwsContainerServiceDeploymentVersion#success_codes}.
	// Experimental.
	SuccessCodes *string `field:"optional" json:"successCodes" yaml:"successCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#timeout_seconds AwsContainerServiceDeploymentVersion#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service_deployment_version#unhealthy_threshold AwsContainerServiceDeploymentVersion#unhealthy_threshold}.
	// Experimental.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

