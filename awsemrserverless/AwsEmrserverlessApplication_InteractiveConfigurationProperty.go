package awsemrserverless


// Experimental.
type AwsEmrserverlessApplication_InteractiveConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#livy_endpoint_enabled AwsEmrserverlessApplication#livy_endpoint_enabled}.
	// Experimental.
	LivyEndpointEnabled interface{} `field:"optional" json:"livyEndpointEnabled" yaml:"livyEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#studio_enabled AwsEmrserverlessApplication#studio_enabled}.
	// Experimental.
	StudioEnabled interface{} `field:"optional" json:"studioEnabled" yaml:"studioEnabled"`
}

