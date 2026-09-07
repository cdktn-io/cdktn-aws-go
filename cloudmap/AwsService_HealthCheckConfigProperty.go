package cloudmap


// Experimental.
type AwsService_HealthCheckConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#failure_threshold AwsService#failure_threshold}.
	// Experimental.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#resource_path AwsService#resource_path}.
	// Experimental.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#type AwsService#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

