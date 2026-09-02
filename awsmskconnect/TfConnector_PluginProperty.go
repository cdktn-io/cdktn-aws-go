package awsmskconnect


// Experimental.
type TfConnector_PluginProperty struct {
	// custom_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#custom_plugin TfConnector#custom_plugin}
	// Experimental.
	CustomPlugin *TfConnector_CustomPluginProperty `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

