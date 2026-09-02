package awsconnect


// Experimental.
type TfQuickConnect_QuickConnectConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#quick_connect_type TfQuickConnect#quick_connect_type}.
	// Experimental.
	QuickConnectType *string `field:"required" json:"quickConnectType" yaml:"quickConnectType"`
	// phone_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#phone_config TfQuickConnect#phone_config}
	// Experimental.
	PhoneConfig interface{} `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// queue_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#queue_config TfQuickConnect#queue_config}
	// Experimental.
	QueueConfig interface{} `field:"optional" json:"queueConfig" yaml:"queueConfig"`
	// user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#user_config TfQuickConnect#user_config}
	// Experimental.
	UserConfig interface{} `field:"optional" json:"userConfig" yaml:"userConfig"`
}

