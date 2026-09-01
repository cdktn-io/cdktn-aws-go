package awselastictranscoder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsElastictranscoderPresetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#container AwsElastictranscoderPreset#container}.
	// Experimental.
	Container *string `field:"required" json:"container" yaml:"container"`
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#audio AwsElastictranscoderPreset#audio}
	// Experimental.
	Audio *AwsElastictranscoderPreset_AudioProperty `field:"optional" json:"audio" yaml:"audio"`
	// audio_codec_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#audio_codec_options AwsElastictranscoderPreset#audio_codec_options}
	// Experimental.
	AudioCodecOptions *AwsElastictranscoderPreset_AudioCodecOptionsProperty `field:"optional" json:"audioCodecOptions" yaml:"audioCodecOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#description AwsElastictranscoderPreset#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#id AwsElastictranscoderPreset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#name AwsElastictranscoderPreset#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#region AwsElastictranscoderPreset#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// thumbnails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#thumbnails AwsElastictranscoderPreset#thumbnails}
	// Experimental.
	Thumbnails *AwsElastictranscoderPreset_ThumbnailsProperty `field:"optional" json:"thumbnails" yaml:"thumbnails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#type AwsElastictranscoderPreset#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#video AwsElastictranscoderPreset#video}
	// Experimental.
	Video *AwsElastictranscoderPreset_VideoProperty `field:"optional" json:"video" yaml:"video"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#video_codec_options AwsElastictranscoderPreset#video_codec_options}.
	// Experimental.
	VideoCodecOptions *map[string]*string `field:"optional" json:"videoCodecOptions" yaml:"videoCodecOptions"`
	// video_watermarks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_preset#video_watermarks AwsElastictranscoderPreset#video_watermarks}
	// Experimental.
	VideoWatermarks interface{} `field:"optional" json:"videoWatermarks" yaml:"videoWatermarks"`
}

