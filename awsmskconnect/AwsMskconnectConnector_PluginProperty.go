package awsmskconnect


// Experimental.
type AwsMskconnectConnector_PluginProperty struct {
	// custom_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#custom_plugin AwsMskconnectConnector#custom_plugin}
	// Experimental.
	CustomPlugin *AwsMskconnectConnector_CustomPluginProperty `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

