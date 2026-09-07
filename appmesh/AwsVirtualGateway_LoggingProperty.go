package appmesh


// Experimental.
type AwsVirtualGateway_LoggingProperty struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#access_log AwsVirtualGateway#access_log}
	// Experimental.
	AccessLog *AwsVirtualGateway_AccessLogProperty `field:"optional" json:"accessLog" yaml:"accessLog"`
}

