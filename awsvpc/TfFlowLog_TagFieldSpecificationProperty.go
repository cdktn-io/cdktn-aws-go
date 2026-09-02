package awsvpc


// Experimental.
type TfFlowLog_TagFieldSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#resource_type TfFlowLog#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#tag_keys TfFlowLog#tag_keys}.
	// Experimental.
	TagKeys *[]*string `field:"required" json:"tagKeys" yaml:"tagKeys"`
}

