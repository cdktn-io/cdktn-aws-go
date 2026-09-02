package awsiotcore


// Experimental.
type TfTopicRule_HttpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#url TfTopicRule#url}.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#confirmation_url TfTopicRule#confirmation_url}.
	// Experimental.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule#http_header TfTopicRule#http_header}
	// Experimental.
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
}

