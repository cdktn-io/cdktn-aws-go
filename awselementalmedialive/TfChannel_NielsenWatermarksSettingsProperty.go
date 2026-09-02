package awselementalmedialive


// Experimental.
type TfChannel_NielsenWatermarksSettingsProperty struct {
	// nielsen_cbet_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_cbet_settings TfChannel#nielsen_cbet_settings}
	// Experimental.
	NielsenCbetSettings *TfChannel_NielsenCbetSettingsProperty `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_distribution_type TfChannel#nielsen_distribution_type}.
	// Experimental.
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// nielsen_naes_ii_nw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_naes_ii_nw_settings TfChannel#nielsen_naes_ii_nw_settings}
	// Experimental.
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

