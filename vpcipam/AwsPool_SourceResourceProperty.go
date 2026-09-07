package vpcipam


// Experimental.
type AwsPool_SourceResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_id AwsPool#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_owner AwsPool#resource_owner}.
	// Experimental.
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_region AwsPool#resource_region}.
	// Experimental.
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_type AwsPool#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

