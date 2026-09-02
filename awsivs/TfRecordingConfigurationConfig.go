package awsivs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRecordingConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#destination_configuration TfRecordingConfiguration#destination_configuration}
	// Experimental.
	DestinationConfiguration *TfRecordingConfiguration_DestinationConfigurationProperty `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#id TfRecordingConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#name TfRecordingConfiguration#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#recording_reconnect_window_seconds TfRecordingConfiguration#recording_reconnect_window_seconds}.
	// Experimental.
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#region TfRecordingConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#tags TfRecordingConfiguration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#tags_all TfRecordingConfiguration#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// thumbnail_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#thumbnail_configuration TfRecordingConfiguration#thumbnail_configuration}
	// Experimental.
	ThumbnailConfiguration *TfRecordingConfiguration_ThumbnailConfigurationProperty `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#timeouts TfRecordingConfiguration#timeouts}
	// Experimental.
	Timeouts *TfRecordingConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

