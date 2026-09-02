package awsvpcipam


// Experimental.
type TfPool_SourceResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_id TfPool#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_owner TfPool#resource_owner}.
	// Experimental.
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_region TfPool#resource_region}.
	// Experimental.
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_type TfPool#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

