package emrserverless


// Experimental.
type AwsApplication_InteractiveConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#livy_endpoint_enabled AwsApplication#livy_endpoint_enabled}.
	// Experimental.
	LivyEndpointEnabled interface{} `field:"optional" json:"livyEndpointEnabled" yaml:"livyEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#studio_enabled AwsApplication#studio_enabled}.
	// Experimental.
	StudioEnabled interface{} `field:"optional" json:"studioEnabled" yaml:"studioEnabled"`
}

