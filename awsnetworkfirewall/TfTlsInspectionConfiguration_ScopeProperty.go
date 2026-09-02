package awsnetworkfirewall


// Experimental.
type TfTlsInspectionConfiguration_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#protocols TfTlsInspectionConfiguration#protocols}.
	// Experimental.
	Protocols *[]*float64 `field:"required" json:"protocols" yaml:"protocols"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#destination TfTlsInspectionConfiguration#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// destination_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#destination_ports TfTlsInspectionConfiguration#destination_ports}
	// Experimental.
	DestinationPorts interface{} `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#source TfTlsInspectionConfiguration#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// source_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#source_ports TfTlsInspectionConfiguration#source_ports}
	// Experimental.
	SourcePorts interface{} `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
}

