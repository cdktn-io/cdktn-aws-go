package awsec2


// Experimental.
type TfLaunchTemplate_InstanceRequirementsProperty struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#memory_mib TfLaunchTemplate#memory_mib}
	// Experimental.
	MemoryMib *TfLaunchTemplate_MemoryMibProperty `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#vcpu_count TfLaunchTemplate#vcpu_count}
	// Experimental.
	VcpuCount *TfLaunchTemplate_VcpuCountProperty `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#accelerator_count TfLaunchTemplate#accelerator_count}
	// Experimental.
	AcceleratorCount *TfLaunchTemplate_AcceleratorCountProperty `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#accelerator_manufacturers TfLaunchTemplate#accelerator_manufacturers}.
	// Experimental.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#accelerator_names TfLaunchTemplate#accelerator_names}.
	// Experimental.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#accelerator_total_memory_mib TfLaunchTemplate#accelerator_total_memory_mib}
	// Experimental.
	AcceleratorTotalMemoryMib *TfLaunchTemplate_AcceleratorTotalMemoryMibProperty `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#accelerator_types TfLaunchTemplate#accelerator_types}.
	// Experimental.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#allowed_instance_types TfLaunchTemplate#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#bare_metal TfLaunchTemplate#bare_metal}.
	// Experimental.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#baseline_ebs_bandwidth_mbps TfLaunchTemplate#baseline_ebs_bandwidth_mbps}
	// Experimental.
	BaselineEbsBandwidthMbps *TfLaunchTemplate_BaselineEbsBandwidthMbpsProperty `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#burstable_performance TfLaunchTemplate#burstable_performance}.
	// Experimental.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#cpu_manufacturers TfLaunchTemplate#cpu_manufacturers}.
	// Experimental.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#excluded_instance_types TfLaunchTemplate#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_generations TfLaunchTemplate#instance_generations}.
	// Experimental.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#local_storage TfLaunchTemplate#local_storage}.
	// Experimental.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#local_storage_types TfLaunchTemplate#local_storage_types}.
	// Experimental.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#max_spot_price_as_percentage_of_optimal_on_demand_price TfLaunchTemplate#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#memory_gib_per_vcpu TfLaunchTemplate#memory_gib_per_vcpu}
	// Experimental.
	MemoryGibPerVcpu *TfLaunchTemplate_MemoryGibPerVcpuProperty `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_bandwidth_gbps TfLaunchTemplate#network_bandwidth_gbps}
	// Experimental.
	NetworkBandwidthGbps *TfLaunchTemplate_NetworkBandwidthGbpsProperty `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_interface_count TfLaunchTemplate#network_interface_count}
	// Experimental.
	NetworkInterfaceCount *TfLaunchTemplate_NetworkInterfaceCountProperty `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#on_demand_max_price_percentage_over_lowest_price TfLaunchTemplate#on_demand_max_price_percentage_over_lowest_price}.
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#require_hibernate_support TfLaunchTemplate#require_hibernate_support}.
	// Experimental.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#spot_max_price_percentage_over_lowest_price TfLaunchTemplate#spot_max_price_percentage_over_lowest_price}.
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#total_local_storage_gb TfLaunchTemplate#total_local_storage_gb}
	// Experimental.
	TotalLocalStorageGb *TfLaunchTemplate_TotalLocalStorageGbProperty `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

