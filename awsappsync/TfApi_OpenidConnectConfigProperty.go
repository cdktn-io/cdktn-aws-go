package awsappsync


// Experimental.
type TfApi_OpenidConnectConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#issuer TfApi#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#auth_ttl TfApi#auth_ttl}.
	// Experimental.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#client_id TfApi#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#iat_ttl TfApi#iat_ttl}.
	// Experimental.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
}

