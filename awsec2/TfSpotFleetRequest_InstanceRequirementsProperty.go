package awsec2


// Experimental.
type TfSpotFleetRequest_InstanceRequirementsProperty struct {
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#accelerator_count TfSpotFleetRequest#accelerator_count}
	// Experimental.
	AcceleratorCount *TfSpotFleetRequest_AcceleratorCountProperty `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#accelerator_manufacturers TfSpotFleetRequest#accelerator_manufacturers}.
	// Experimental.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#accelerator_names TfSpotFleetRequest#accelerator_names}.
	// Experimental.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#accelerator_total_memory_mib TfSpotFleetRequest#accelerator_total_memory_mib}
	// Experimental.
	AcceleratorTotalMemoryMib *TfSpotFleetRequest_AcceleratorTotalMemoryMibProperty `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#accelerator_types TfSpotFleetRequest#accelerator_types}.
	// Experimental.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#allowed_instance_types TfSpotFleetRequest#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#bare_metal TfSpotFleetRequest#bare_metal}.
	// Experimental.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#baseline_ebs_bandwidth_mbps TfSpotFleetRequest#baseline_ebs_bandwidth_mbps}
	// Experimental.
	BaselineEbsBandwidthMbps *TfSpotFleetRequest_BaselineEbsBandwidthMbpsProperty `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#burstable_performance TfSpotFleetRequest#burstable_performance}.
	// Experimental.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#cpu_manufacturers TfSpotFleetRequest#cpu_manufacturers}.
	// Experimental.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#excluded_instance_types TfSpotFleetRequest#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#instance_generations TfSpotFleetRequest#instance_generations}.
	// Experimental.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#local_storage TfSpotFleetRequest#local_storage}.
	// Experimental.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#local_storage_types TfSpotFleetRequest#local_storage_types}.
	// Experimental.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#memory_gib_per_vcpu TfSpotFleetRequest#memory_gib_per_vcpu}
	// Experimental.
	MemoryGibPerVcpu *TfSpotFleetRequest_MemoryGibPerVcpuProperty `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#memory_mib TfSpotFleetRequest#memory_mib}
	// Experimental.
	MemoryMib *TfSpotFleetRequest_MemoryMibProperty `field:"optional" json:"memoryMib" yaml:"memoryMib"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#network_bandwidth_gbps TfSpotFleetRequest#network_bandwidth_gbps}
	// Experimental.
	NetworkBandwidthGbps *TfSpotFleetRequest_NetworkBandwidthGbpsProperty `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#network_interface_count TfSpotFleetRequest#network_interface_count}
	// Experimental.
	NetworkInterfaceCount *TfSpotFleetRequest_NetworkInterfaceCountProperty `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#on_demand_max_price_percentage_over_lowest_price TfSpotFleetRequest#on_demand_max_price_percentage_over_lowest_price}.
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#require_hibernate_support TfSpotFleetRequest#require_hibernate_support}.
	// Experimental.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#spot_max_price_percentage_over_lowest_price TfSpotFleetRequest#spot_max_price_percentage_over_lowest_price}.
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#total_local_storage_gb TfSpotFleetRequest#total_local_storage_gb}
	// Experimental.
	TotalLocalStorageGb *TfSpotFleetRequest_TotalLocalStorageGbProperty `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#vcpu_count TfSpotFleetRequest#vcpu_count}
	// Experimental.
	VcpuCount *TfSpotFleetRequest_VcpuCountProperty `field:"optional" json:"vcpuCount" yaml:"vcpuCount"`
}

