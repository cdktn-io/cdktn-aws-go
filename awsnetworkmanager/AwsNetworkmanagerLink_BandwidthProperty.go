package awsnetworkmanager


// Experimental.
type AwsNetworkmanagerLink_BandwidthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_link#download_speed AwsNetworkmanagerLink#download_speed}.
	// Experimental.
	DownloadSpeed *float64 `field:"optional" json:"downloadSpeed" yaml:"downloadSpeed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_link#upload_speed AwsNetworkmanagerLink#upload_speed}.
	// Experimental.
	UploadSpeed *float64 `field:"optional" json:"uploadSpeed" yaml:"uploadSpeed"`
}

