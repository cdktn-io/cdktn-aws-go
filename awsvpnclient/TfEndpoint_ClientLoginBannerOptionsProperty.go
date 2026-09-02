package awsvpnclient


// Experimental.
type TfEndpoint_ClientLoginBannerOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#banner_text TfEndpoint#banner_text}.
	// Experimental.
	BannerText *string `field:"optional" json:"bannerText" yaml:"bannerText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#enabled TfEndpoint#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

