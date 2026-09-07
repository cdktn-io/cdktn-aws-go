package s3


// Experimental.
type AwsBucket_WebsiteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#error_document AwsBucket#error_document}.
	// Experimental.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#index_document AwsBucket#index_document}.
	// Experimental.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#redirect_all_requests_to AwsBucket#redirect_all_requests_to}.
	// Experimental.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#routing_rules AwsBucket#routing_rules}.
	// Experimental.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}

