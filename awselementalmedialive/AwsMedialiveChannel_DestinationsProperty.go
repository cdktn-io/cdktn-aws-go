package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_DestinationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#id AwsMedialiveChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// media_package_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_settings AwsMedialiveChannel#media_package_settings}
	// Experimental.
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// multiplex_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_settings AwsMedialiveChannel#multiplex_settings}
	// Experimental.
	MultiplexSettings *AwsMedialiveChannel_MultiplexSettingsProperty `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#settings AwsMedialiveChannel#settings}
	// Experimental.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

