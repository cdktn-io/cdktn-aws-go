package awscloudfront


// Experimental.
type TfDistribution_OriginGroupProperty struct {
	// failover_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#failover_criteria TfDistribution#failover_criteria}
	// Experimental.
	FailoverCriteria *TfDistribution_FailoverCriteriaProperty `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#member TfDistribution#member}
	// Experimental.
	Member interface{} `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_id TfDistribution#origin_id}.
	// Experimental.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

