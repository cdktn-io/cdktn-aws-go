package awsworkspacesweb


// Experimental.
type TfSessionLogger_EventFilterProperty struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#all TfSessionLogger#all}
	// Experimental.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#include TfSessionLogger#include}.
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

