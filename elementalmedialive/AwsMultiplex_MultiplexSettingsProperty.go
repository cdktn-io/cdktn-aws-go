package elementalmedialive


// Experimental.
type AwsMultiplex_MultiplexSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex#transport_stream_bitrate AwsMultiplex#transport_stream_bitrate}.
	// Experimental.
	TransportStreamBitrate *float64 `field:"required" json:"transportStreamBitrate" yaml:"transportStreamBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex#transport_stream_id AwsMultiplex#transport_stream_id}.
	// Experimental.
	TransportStreamId *float64 `field:"required" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex#maximum_video_buffer_delay_milliseconds AwsMultiplex#maximum_video_buffer_delay_milliseconds}.
	// Experimental.
	MaximumVideoBufferDelayMilliseconds *float64 `field:"optional" json:"maximumVideoBufferDelayMilliseconds" yaml:"maximumVideoBufferDelayMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex#transport_stream_reserved_bitrate AwsMultiplex#transport_stream_reserved_bitrate}.
	// Experimental.
	TransportStreamReservedBitrate *float64 `field:"optional" json:"transportStreamReservedBitrate" yaml:"transportStreamReservedBitrate"`
}

