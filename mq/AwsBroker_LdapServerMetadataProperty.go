package mq


// Experimental.
type AwsBroker_LdapServerMetadataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#hosts AwsBroker#hosts}.
	// Experimental.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_base AwsBroker#role_base}.
	// Experimental.
	RoleBase *string `field:"optional" json:"roleBase" yaml:"roleBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_name AwsBroker#role_name}.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_search_matching AwsBroker#role_search_matching}.
	// Experimental.
	RoleSearchMatching *string `field:"optional" json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#role_search_subtree AwsBroker#role_search_subtree}.
	// Experimental.
	RoleSearchSubtree interface{} `field:"optional" json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#service_account_password AwsBroker#service_account_password}.
	// Experimental.
	ServiceAccountPassword *string `field:"optional" json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#service_account_username AwsBroker#service_account_username}.
	// Experimental.
	ServiceAccountUsername *string `field:"optional" json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_base AwsBroker#user_base}.
	// Experimental.
	UserBase *string `field:"optional" json:"userBase" yaml:"userBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_role_name AwsBroker#user_role_name}.
	// Experimental.
	UserRoleName *string `field:"optional" json:"userRoleName" yaml:"userRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_search_matching AwsBroker#user_search_matching}.
	// Experimental.
	UserSearchMatching *string `field:"optional" json:"userSearchMatching" yaml:"userSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#user_search_subtree AwsBroker#user_search_subtree}.
	// Experimental.
	UserSearchSubtree interface{} `field:"optional" json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

