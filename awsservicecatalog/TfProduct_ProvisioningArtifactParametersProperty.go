package awsservicecatalog


// Experimental.
type TfProduct_ProvisioningArtifactParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#description TfProduct#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#disable_template_validation TfProduct#disable_template_validation}.
	// Experimental.
	DisableTemplateValidation interface{} `field:"optional" json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#name TfProduct#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#template_physical_id TfProduct#template_physical_id}.
	// Experimental.
	TemplatePhysicalId *string `field:"optional" json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#template_url TfProduct#template_url}.
	// Experimental.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_product#type TfProduct#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

