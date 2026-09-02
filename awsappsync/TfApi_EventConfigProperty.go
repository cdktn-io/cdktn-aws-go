package awsappsync


// Experimental.
type TfApi_EventConfigProperty struct {
	// auth_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#auth_provider TfApi#auth_provider}
	// Experimental.
	AuthProvider interface{} `field:"optional" json:"authProvider" yaml:"authProvider"`
	// connection_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#connection_auth_mode TfApi#connection_auth_mode}
	// Experimental.
	ConnectionAuthMode interface{} `field:"optional" json:"connectionAuthMode" yaml:"connectionAuthMode"`
	// default_publish_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#default_publish_auth_mode TfApi#default_publish_auth_mode}
	// Experimental.
	DefaultPublishAuthMode interface{} `field:"optional" json:"defaultPublishAuthMode" yaml:"defaultPublishAuthMode"`
	// default_subscribe_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#default_subscribe_auth_mode TfApi#default_subscribe_auth_mode}
	// Experimental.
	DefaultSubscribeAuthMode interface{} `field:"optional" json:"defaultSubscribeAuthMode" yaml:"defaultSubscribeAuthMode"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#log_config TfApi#log_config}
	// Experimental.
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
}

