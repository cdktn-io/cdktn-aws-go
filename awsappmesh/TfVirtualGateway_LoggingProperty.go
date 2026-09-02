package awsappmesh


// Experimental.
type TfVirtualGateway_LoggingProperty struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#access_log TfVirtualGateway#access_log}
	// Experimental.
	AccessLog *TfVirtualGateway_AccessLogProperty `field:"optional" json:"accessLog" yaml:"accessLog"`
}

