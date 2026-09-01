package awstransferfamily


// Experimental.
type AwsTransferServer_ProtocolDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#as2_transports AwsTransferServer#as2_transports}.
	// Experimental.
	As2Transports *[]*string `field:"optional" json:"as2Transports" yaml:"as2Transports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#passive_ip AwsTransferServer#passive_ip}.
	// Experimental.
	PassiveIp *string `field:"optional" json:"passiveIp" yaml:"passiveIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#set_stat_option AwsTransferServer#set_stat_option}.
	// Experimental.
	SetStatOption *string `field:"optional" json:"setStatOption" yaml:"setStatOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#tls_session_resumption_mode AwsTransferServer#tls_session_resumption_mode}.
	// Experimental.
	TlsSessionResumptionMode *string `field:"optional" json:"tlsSessionResumptionMode" yaml:"tlsSessionResumptionMode"`
}

