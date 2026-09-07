package route53


// Experimental.
type AwsRecord_CidrRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#collection_id AwsRecord#collection_id}.
	// Experimental.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#location_name AwsRecord#location_name}.
	// Experimental.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}

