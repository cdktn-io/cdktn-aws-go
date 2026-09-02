package awsopensearch


// Experimental.
type TfDomain_LogPublishingOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#cloudwatch_log_group_arn TfDomain#cloudwatch_log_group_arn}.
	// Experimental.
	CloudwatchLogGroupArn *string `field:"required" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#log_type TfDomain#log_type}.
	// Experimental.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

