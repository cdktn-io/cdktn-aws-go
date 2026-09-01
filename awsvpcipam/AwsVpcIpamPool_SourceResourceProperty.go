package awsvpcipam


// Experimental.
type AwsVpcIpamPool_SourceResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_id AwsVpcIpamPool#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_owner AwsVpcIpamPool#resource_owner}.
	// Experimental.
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_region AwsVpcIpamPool#resource_region}.
	// Experimental.
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#resource_type AwsVpcIpamPool#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

