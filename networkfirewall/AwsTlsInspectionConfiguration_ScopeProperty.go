package networkfirewall


// Experimental.
type AwsTlsInspectionConfiguration_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#protocols AwsTlsInspectionConfiguration#protocols}.
	// Experimental.
	Protocols *[]*float64 `field:"required" json:"protocols" yaml:"protocols"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#destination AwsTlsInspectionConfiguration#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// destination_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#destination_ports AwsTlsInspectionConfiguration#destination_ports}
	// Experimental.
	DestinationPorts interface{} `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#source AwsTlsInspectionConfiguration#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// source_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#source_ports AwsTlsInspectionConfiguration#source_ports}
	// Experimental.
	SourcePorts interface{} `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
}

