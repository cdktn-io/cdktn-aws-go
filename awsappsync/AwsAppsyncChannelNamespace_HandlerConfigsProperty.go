package awsappsync


// Experimental.
type AwsAppsyncChannelNamespace_HandlerConfigsProperty struct {
	// on_publish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#on_publish AwsAppsyncChannelNamespace#on_publish}
	// Experimental.
	OnPublish interface{} `field:"optional" json:"onPublish" yaml:"onPublish"`
	// on_subscribe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#on_subscribe AwsAppsyncChannelNamespace#on_subscribe}
	// Experimental.
	OnSubscribe interface{} `field:"optional" json:"onSubscribe" yaml:"onSubscribe"`
}

