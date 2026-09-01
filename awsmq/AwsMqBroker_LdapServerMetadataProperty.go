package awsmq


// Experimental.
type AwsMqBroker_LdapServerMetadataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#hosts AwsMqBroker#hosts}.
	// Experimental.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_base AwsMqBroker#role_base}.
	// Experimental.
	RoleBase *string `field:"optional" json:"roleBase" yaml:"roleBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_name AwsMqBroker#role_name}.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_search_matching AwsMqBroker#role_search_matching}.
	// Experimental.
	RoleSearchMatching *string `field:"optional" json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_search_subtree AwsMqBroker#role_search_subtree}.
	// Experimental.
	RoleSearchSubtree interface{} `field:"optional" json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#service_account_password AwsMqBroker#service_account_password}.
	// Experimental.
	ServiceAccountPassword *string `field:"optional" json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#service_account_username AwsMqBroker#service_account_username}.
	// Experimental.
	ServiceAccountUsername *string `field:"optional" json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_base AwsMqBroker#user_base}.
	// Experimental.
	UserBase *string `field:"optional" json:"userBase" yaml:"userBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_role_name AwsMqBroker#user_role_name}.
	// Experimental.
	UserRoleName *string `field:"optional" json:"userRoleName" yaml:"userRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_search_matching AwsMqBroker#user_search_matching}.
	// Experimental.
	UserSearchMatching *string `field:"optional" json:"userSearchMatching" yaml:"userSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_search_subtree AwsMqBroker#user_search_subtree}.
	// Experimental.
	UserSearchSubtree interface{} `field:"optional" json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

