package awselementalmedialive


// Experimental.
type TfChannel_M3u8SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_frames_per_pes TfChannel#audio_frames_per_pes}.
	// Experimental.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_pids TfChannel#audio_pids}.
	// Experimental.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ecm_pid TfChannel#ecm_pid}.
	// Experimental.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_id3_behavior TfChannel#nielsen_id3_behavior}.
	// Experimental.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pat_interval TfChannel#pat_interval}.
	// Experimental.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_control TfChannel#pcr_control}.
	// Experimental.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_period TfChannel#pcr_period}.
	// Experimental.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_pid TfChannel#pcr_pid}.
	// Experimental.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pmt_interval TfChannel#pmt_interval}.
	// Experimental.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pmt_pid TfChannel#pmt_pid}.
	// Experimental.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_num TfChannel#program_num}.
	// Experimental.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_behavior TfChannel#scte35_behavior}.
	// Experimental.
	Scte35Behavior *string `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_pid TfChannel#scte35_pid}.
	// Experimental.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_behavior TfChannel#timed_metadata_behavior}.
	// Experimental.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_pid TfChannel#timed_metadata_pid}.
	// Experimental.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#transport_stream_id TfChannel#transport_stream_id}.
	// Experimental.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_pid TfChannel#video_pid}.
	// Experimental.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

