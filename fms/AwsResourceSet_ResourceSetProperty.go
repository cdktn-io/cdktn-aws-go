package fms


// Experimental.
type AwsResourceSet_ResourceSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_resource_set#name AwsResourceSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_resource_set#description AwsResourceSet#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_resource_set#resource_set_status AwsResourceSet#resource_set_status}.
	// Experimental.
	ResourceSetStatus *string `field:"optional" json:"resourceSetStatus" yaml:"resourceSetStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_resource_set#resource_type_list AwsResourceSet#resource_type_list}.
	// Experimental.
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_resource_set#update_token AwsResourceSet#update_token}.
	// Experimental.
	UpdateToken *string `field:"optional" json:"updateToken" yaml:"updateToken"`
}

