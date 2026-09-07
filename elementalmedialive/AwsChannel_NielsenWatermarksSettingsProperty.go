package elementalmedialive


// Experimental.
type AwsChannel_NielsenWatermarksSettingsProperty struct {
	// nielsen_cbet_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_cbet_settings AwsChannel#nielsen_cbet_settings}
	// Experimental.
	NielsenCbetSettings *AwsChannel_NielsenCbetSettingsProperty `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_distribution_type AwsChannel#nielsen_distribution_type}.
	// Experimental.
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// nielsen_naes_ii_nw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_naes_ii_nw_settings AwsChannel#nielsen_naes_ii_nw_settings}
	// Experimental.
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

