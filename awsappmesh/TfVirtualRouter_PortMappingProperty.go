package awsappmesh


// Experimental.
type TfVirtualRouter_PortMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_router#port TfVirtualRouter#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_router#protocol TfVirtualRouter#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

