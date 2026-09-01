package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_LoggingProperty struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#access_log AwsAppmeshVirtualGateway#access_log}
	// Experimental.
	AccessLog *AwsAppmeshVirtualGateway_AccessLogProperty `field:"optional" json:"accessLog" yaml:"accessLog"`
}

