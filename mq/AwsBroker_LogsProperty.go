package mq


// Experimental.
type AwsBroker_LogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#audit AwsBroker#audit}.
	// Experimental.
	Audit *string `field:"optional" json:"audit" yaml:"audit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#general AwsBroker#general}.
	// Experimental.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

