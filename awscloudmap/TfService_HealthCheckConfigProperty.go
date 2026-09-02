package awscloudmap


// Experimental.
type TfService_HealthCheckConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#failure_threshold TfService#failure_threshold}.
	// Experimental.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#resource_path TfService#resource_path}.
	// Experimental.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/service_discovery_service#type TfService#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

