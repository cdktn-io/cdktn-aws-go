package appflow


// Experimental.
type AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#entity_name AwsFlow#entity_name}.
	// Experimental.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_properties AwsFlow#custom_properties}.
	// Experimental.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
}

