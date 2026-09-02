package awsappflow


// Experimental.
type TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#entity_name TfFlow#entity_name}.
	// Experimental.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_properties TfFlow#custom_properties}.
	// Experimental.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
}

