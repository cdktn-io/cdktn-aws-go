package awslexv2models


// Experimental.
type TfIntent_OpensearchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#domain_endpoint TfIntent#domain_endpoint}.
	// Experimental.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#index_name TfIntent#index_name}.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response TfIntent#exact_response}.
	// Experimental.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// exact_response_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response_fields TfIntent#exact_response_fields}
	// Experimental.
	ExactResponseFields interface{} `field:"optional" json:"exactResponseFields" yaml:"exactResponseFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#include_fields TfIntent#include_fields}.
	// Experimental.
	IncludeFields *[]*string `field:"optional" json:"includeFields" yaml:"includeFields"`
}

