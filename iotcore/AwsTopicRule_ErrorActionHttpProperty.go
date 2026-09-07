package iotcore


// Experimental.
type AwsTopicRule_ErrorActionHttpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#url AwsTopicRule#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#confirmation_url AwsTopicRule#confirmation_url}.
	// Experimental.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http_header AwsTopicRule#http_header}
	// Experimental.
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
}

