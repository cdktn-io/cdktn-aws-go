package glue


// Experimental.
type AwsUserDefinedFunction_ResourceUrisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_user_defined_function#resource_type AwsUserDefinedFunction#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_user_defined_function#uri AwsUserDefinedFunction#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

