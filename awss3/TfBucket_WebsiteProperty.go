package awss3


// Experimental.
type TfBucket_WebsiteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#error_document TfBucket#error_document}.
	// Experimental.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#index_document TfBucket#index_document}.
	// Experimental.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#redirect_all_requests_to TfBucket#redirect_all_requests_to}.
	// Experimental.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#routing_rules TfBucket#routing_rules}.
	// Experimental.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}

