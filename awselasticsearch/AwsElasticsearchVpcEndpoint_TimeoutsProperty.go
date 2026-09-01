package awselasticsearch


// Experimental.
type AwsElasticsearchVpcEndpoint_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_vpc_endpoint#create AwsElasticsearchVpcEndpoint#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_vpc_endpoint#delete AwsElasticsearchVpcEndpoint#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_vpc_endpoint#update AwsElasticsearchVpcEndpoint#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

