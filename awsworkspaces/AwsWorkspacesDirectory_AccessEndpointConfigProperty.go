package awsworkspaces


// Experimental.
type AwsWorkspacesDirectory_AccessEndpointConfigProperty struct {
	// access_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#access_endpoints AwsWorkspacesDirectory#access_endpoints}
	// Experimental.
	AccessEndpoints interface{} `field:"required" json:"accessEndpoints" yaml:"accessEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#internet_fallback_protocols AwsWorkspacesDirectory#internet_fallback_protocols}.
	// Experimental.
	InternetFallbackProtocols *[]*string `field:"optional" json:"internetFallbackProtocols" yaml:"internetFallbackProtocols"`
}

