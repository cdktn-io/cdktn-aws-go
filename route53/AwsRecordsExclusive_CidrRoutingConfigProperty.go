package route53


// Experimental.
type AwsRecordsExclusive_CidrRoutingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#collection_id AwsRecordsExclusive#collection_id}.
	// Experimental.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#location_name AwsRecordsExclusive#location_name}.
	// Experimental.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}

