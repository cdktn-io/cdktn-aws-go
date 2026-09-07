package sesmailmanager


// Experimental.
type AwsRuleSet_ActionProperty struct {
	// add_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#add_header AwsRuleSet#add_header}
	// Experimental.
	AddHeader interface{} `field:"optional" json:"addHeader" yaml:"addHeader"`
	// archive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#archive AwsRuleSet#archive}
	// Experimental.
	Archive interface{} `field:"optional" json:"archive" yaml:"archive"`
	// bounce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#bounce AwsRuleSet#bounce}
	// Experimental.
	Bounce interface{} `field:"optional" json:"bounce" yaml:"bounce"`
	// deliver_to_mailbox block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#deliver_to_mailbox AwsRuleSet#deliver_to_mailbox}
	// Experimental.
	DeliverToMailbox interface{} `field:"optional" json:"deliverToMailbox" yaml:"deliverToMailbox"`
	// deliver_to_q_business block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#deliver_to_q_business AwsRuleSet#deliver_to_q_business}
	// Experimental.
	DeliverToQBusiness interface{} `field:"optional" json:"deliverToQBusiness" yaml:"deliverToQBusiness"`
	// drop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#drop AwsRuleSet#drop}
	// Experimental.
	Drop interface{} `field:"optional" json:"drop" yaml:"drop"`
	// invoke_lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#invoke_lambda AwsRuleSet#invoke_lambda}
	// Experimental.
	InvokeLambda interface{} `field:"optional" json:"invokeLambda" yaml:"invokeLambda"`
	// publish_to_sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#publish_to_sns AwsRuleSet#publish_to_sns}
	// Experimental.
	PublishToSns interface{} `field:"optional" json:"publishToSns" yaml:"publishToSns"`
	// relay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#relay AwsRuleSet#relay}
	// Experimental.
	Relay interface{} `field:"optional" json:"relay" yaml:"relay"`
	// replace_recipient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#replace_recipient AwsRuleSet#replace_recipient}
	// Experimental.
	ReplaceRecipient interface{} `field:"optional" json:"replaceRecipient" yaml:"replaceRecipient"`
	// send block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#send AwsRuleSet#send}
	// Experimental.
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// write_to_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#write_to_s3 AwsRuleSet#write_to_s3}
	// Experimental.
	WriteToS3 interface{} `field:"optional" json:"writeToS3" yaml:"writeToS3"`
}

