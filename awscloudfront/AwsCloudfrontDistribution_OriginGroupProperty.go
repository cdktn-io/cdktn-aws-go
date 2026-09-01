package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_OriginGroupProperty struct {
	// failover_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#failover_criteria AwsCloudfrontDistribution#failover_criteria}
	// Experimental.
	FailoverCriteria *AwsCloudfrontDistribution_FailoverCriteriaProperty `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#member AwsCloudfrontDistribution#member}
	// Experimental.
	Member interface{} `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_id AwsCloudfrontDistribution#origin_id}.
	// Experimental.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

