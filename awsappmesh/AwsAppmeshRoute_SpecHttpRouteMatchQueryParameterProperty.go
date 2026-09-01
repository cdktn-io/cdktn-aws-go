package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecHttpRouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#name AwsAppmeshRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsAppmeshRoute#match}
	// Experimental.
	Match *AwsAppmeshRoute_SpecHttpRouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

