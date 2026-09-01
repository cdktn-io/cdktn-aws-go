package awsmq


// Experimental.
type AwsMqBroker_LogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#audit AwsMqBroker#audit}.
	// Experimental.
	Audit *string `field:"optional" json:"audit" yaml:"audit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#general AwsMqBroker#general}.
	// Experimental.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

