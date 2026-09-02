package awsmq


// Experimental.
type TfBroker_LogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#audit TfBroker#audit}.
	// Experimental.
	Audit *string `field:"optional" json:"audit" yaml:"audit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#general TfBroker#general}.
	// Experimental.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

