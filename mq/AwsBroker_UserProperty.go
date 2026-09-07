package mq


// Experimental.
type AwsBroker_UserProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#password AwsBroker#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#username AwsBroker#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#console_access AwsBroker#console_access}.
	// Experimental.
	ConsoleAccess interface{} `field:"optional" json:"consoleAccess" yaml:"consoleAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#groups AwsBroker#groups}.
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#replication_user AwsBroker#replication_user}.
	// Experimental.
	ReplicationUser interface{} `field:"optional" json:"replicationUser" yaml:"replicationUser"`
}

