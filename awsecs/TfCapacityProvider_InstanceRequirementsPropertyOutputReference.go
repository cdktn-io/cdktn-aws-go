package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCapacityProvider_InstanceRequirementsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcceleratorCount() TfCapacityProvider_AcceleratorCountPropertyOutputReference
	// Experimental.
	AcceleratorCountInput() *TfCapacityProvider_AcceleratorCountProperty
	// Experimental.
	AcceleratorManufacturers() *[]*string
	// Experimental.
	SetAcceleratorManufacturers(val *[]*string)
	// Experimental.
	AcceleratorManufacturersInput() *[]*string
	// Experimental.
	AcceleratorNames() *[]*string
	// Experimental.
	SetAcceleratorNames(val *[]*string)
	// Experimental.
	AcceleratorNamesInput() *[]*string
	// Experimental.
	AcceleratorTotalMemoryMib() TfCapacityProvider_AcceleratorTotalMemoryMibPropertyOutputReference
	// Experimental.
	AcceleratorTotalMemoryMibInput() *TfCapacityProvider_AcceleratorTotalMemoryMibProperty
	// Experimental.
	AcceleratorTypes() *[]*string
	// Experimental.
	SetAcceleratorTypes(val *[]*string)
	// Experimental.
	AcceleratorTypesInput() *[]*string
	// Experimental.
	AllowedInstanceTypes() *[]*string
	// Experimental.
	SetAllowedInstanceTypes(val *[]*string)
	// Experimental.
	AllowedInstanceTypesInput() *[]*string
	// Experimental.
	BareMetal() *string
	// Experimental.
	SetBareMetal(val *string)
	// Experimental.
	BareMetalInput() *string
	// Experimental.
	BaselineEbsBandwidthMbps() TfCapacityProvider_BaselineEbsBandwidthMbpsPropertyOutputReference
	// Experimental.
	BaselineEbsBandwidthMbpsInput() *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty
	// Experimental.
	BurstablePerformance() *string
	// Experimental.
	SetBurstablePerformance(val *string)
	// Experimental.
	BurstablePerformanceInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	CpuManufacturers() *[]*string
	// Experimental.
	SetCpuManufacturers(val *[]*string)
	// Experimental.
	CpuManufacturersInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ExcludedInstanceTypes() *[]*string
	// Experimental.
	SetExcludedInstanceTypes(val *[]*string)
	// Experimental.
	ExcludedInstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceGenerations() *[]*string
	// Experimental.
	SetInstanceGenerations(val *[]*string)
	// Experimental.
	InstanceGenerationsInput() *[]*string
	// Experimental.
	InternalValue() *TfCapacityProvider_InstanceRequirementsProperty
	// Experimental.
	SetInternalValue(val *TfCapacityProvider_InstanceRequirementsProperty)
	// Experimental.
	LocalStorage() *string
	// Experimental.
	SetLocalStorage(val *string)
	// Experimental.
	LocalStorageInput() *string
	// Experimental.
	LocalStorageTypes() *[]*string
	// Experimental.
	SetLocalStorageTypes(val *[]*string)
	// Experimental.
	LocalStorageTypesInput() *[]*string
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	// Experimental.
	SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64)
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64
	// Experimental.
	MemoryGibPerVcpu() TfCapacityProvider_MemoryGibPerVcpuPropertyOutputReference
	// Experimental.
	MemoryGibPerVcpuInput() *TfCapacityProvider_MemoryGibPerVcpuProperty
	// Experimental.
	MemoryMib() TfCapacityProvider_MemoryMibPropertyOutputReference
	// Experimental.
	MemoryMibInput() *TfCapacityProvider_MemoryMibProperty
	// Experimental.
	NetworkBandwidthGbps() TfCapacityProvider_NetworkBandwidthGbpsPropertyOutputReference
	// Experimental.
	NetworkBandwidthGbpsInput() *TfCapacityProvider_NetworkBandwidthGbpsProperty
	// Experimental.
	NetworkInterfaceCount() TfCapacityProvider_NetworkInterfaceCountPropertyOutputReference
	// Experimental.
	NetworkInterfaceCountInput() *TfCapacityProvider_NetworkInterfaceCountProperty
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice() *float64
	// Experimental.
	SetOnDemandMaxPricePercentageOverLowestPrice(val *float64)
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPriceInput() *float64
	// Experimental.
	RequireHibernateSupport() interface{}
	// Experimental.
	SetRequireHibernateSupport(val interface{})
	// Experimental.
	RequireHibernateSupportInput() interface{}
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice() *float64
	// Experimental.
	SetSpotMaxPricePercentageOverLowestPrice(val *float64)
	// Experimental.
	SpotMaxPricePercentageOverLowestPriceInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TotalLocalStorageGb() TfCapacityProvider_TotalLocalStorageGbPropertyOutputReference
	// Experimental.
	TotalLocalStorageGbInput() *TfCapacityProvider_TotalLocalStorageGbProperty
	// Experimental.
	VcpuCount() TfCapacityProvider_VcpuCountPropertyOutputReference
	// Experimental.
	VcpuCountInput() *TfCapacityProvider_VcpuCountProperty
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAcceleratorCount(value *TfCapacityProvider_AcceleratorCountProperty)
	// Experimental.
	PutAcceleratorTotalMemoryMib(value *TfCapacityProvider_AcceleratorTotalMemoryMibProperty)
	// Experimental.
	PutBaselineEbsBandwidthMbps(value *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty)
	// Experimental.
	PutMemoryGibPerVcpu(value *TfCapacityProvider_MemoryGibPerVcpuProperty)
	// Experimental.
	PutMemoryMib(value *TfCapacityProvider_MemoryMibProperty)
	// Experimental.
	PutNetworkBandwidthGbps(value *TfCapacityProvider_NetworkBandwidthGbpsProperty)
	// Experimental.
	PutNetworkInterfaceCount(value *TfCapacityProvider_NetworkInterfaceCountProperty)
	// Experimental.
	PutTotalLocalStorageGb(value *TfCapacityProvider_TotalLocalStorageGbProperty)
	// Experimental.
	PutVcpuCount(value *TfCapacityProvider_VcpuCountProperty)
	// Experimental.
	ResetAcceleratorCount()
	// Experimental.
	ResetAcceleratorManufacturers()
	// Experimental.
	ResetAcceleratorNames()
	// Experimental.
	ResetAcceleratorTotalMemoryMib()
	// Experimental.
	ResetAcceleratorTypes()
	// Experimental.
	ResetAllowedInstanceTypes()
	// Experimental.
	ResetBareMetal()
	// Experimental.
	ResetBaselineEbsBandwidthMbps()
	// Experimental.
	ResetBurstablePerformance()
	// Experimental.
	ResetCpuManufacturers()
	// Experimental.
	ResetExcludedInstanceTypes()
	// Experimental.
	ResetInstanceGenerations()
	// Experimental.
	ResetLocalStorage()
	// Experimental.
	ResetLocalStorageTypes()
	// Experimental.
	ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice()
	// Experimental.
	ResetMemoryGibPerVcpu()
	// Experimental.
	ResetNetworkBandwidthGbps()
	// Experimental.
	ResetNetworkInterfaceCount()
	// Experimental.
	ResetOnDemandMaxPricePercentageOverLowestPrice()
	// Experimental.
	ResetRequireHibernateSupport()
	// Experimental.
	ResetSpotMaxPricePercentageOverLowestPrice()
	// Experimental.
	ResetTotalLocalStorageGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCapacityProvider_InstanceRequirementsPropertyOutputReference
type jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorCount() TfCapacityProvider_AcceleratorCountPropertyOutputReference {
	var returns TfCapacityProvider_AcceleratorCountPropertyOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorCountInput() *TfCapacityProvider_AcceleratorCountProperty {
	var returns *TfCapacityProvider_AcceleratorCountProperty
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorTotalMemoryMib() TfCapacityProvider_AcceleratorTotalMemoryMibPropertyOutputReference {
	var returns TfCapacityProvider_AcceleratorTotalMemoryMibPropertyOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorTotalMemoryMibInput() *TfCapacityProvider_AcceleratorTotalMemoryMibProperty {
	var returns *TfCapacityProvider_AcceleratorTotalMemoryMibProperty
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) AllowedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BaselineEbsBandwidthMbps() TfCapacityProvider_BaselineEbsBandwidthMbpsPropertyOutputReference {
	var returns TfCapacityProvider_BaselineEbsBandwidthMbpsPropertyOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BaselineEbsBandwidthMbpsInput() *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty {
	var returns *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) InternalValue() *TfCapacityProvider_InstanceRequirementsProperty {
	var returns *TfCapacityProvider_InstanceRequirementsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MemoryGibPerVcpu() TfCapacityProvider_MemoryGibPerVcpuPropertyOutputReference {
	var returns TfCapacityProvider_MemoryGibPerVcpuPropertyOutputReference
	_jsii_.Get(
		j,
		"memoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MemoryGibPerVcpuInput() *TfCapacityProvider_MemoryGibPerVcpuProperty {
	var returns *TfCapacityProvider_MemoryGibPerVcpuProperty
	_jsii_.Get(
		j,
		"memoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MemoryMib() TfCapacityProvider_MemoryMibPropertyOutputReference {
	var returns TfCapacityProvider_MemoryMibPropertyOutputReference
	_jsii_.Get(
		j,
		"memoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) MemoryMibInput() *TfCapacityProvider_MemoryMibProperty {
	var returns *TfCapacityProvider_MemoryMibProperty
	_jsii_.Get(
		j,
		"memoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) NetworkBandwidthGbps() TfCapacityProvider_NetworkBandwidthGbpsPropertyOutputReference {
	var returns TfCapacityProvider_NetworkBandwidthGbpsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) NetworkBandwidthGbpsInput() *TfCapacityProvider_NetworkBandwidthGbpsProperty {
	var returns *TfCapacityProvider_NetworkBandwidthGbpsProperty
	_jsii_.Get(
		j,
		"networkBandwidthGbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) NetworkInterfaceCount() TfCapacityProvider_NetworkInterfaceCountPropertyOutputReference {
	var returns TfCapacityProvider_NetworkInterfaceCountPropertyOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) NetworkInterfaceCountInput() *TfCapacityProvider_NetworkInterfaceCountProperty {
	var returns *TfCapacityProvider_NetworkInterfaceCountProperty
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) TotalLocalStorageGb() TfCapacityProvider_TotalLocalStorageGbPropertyOutputReference {
	var returns TfCapacityProvider_TotalLocalStorageGbPropertyOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) TotalLocalStorageGbInput() *TfCapacityProvider_TotalLocalStorageGbProperty {
	var returns *TfCapacityProvider_TotalLocalStorageGbProperty
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) VcpuCount() TfCapacityProvider_VcpuCountPropertyOutputReference {
	var returns TfCapacityProvider_VcpuCountPropertyOutputReference
	_jsii_.Get(
		j,
		"vcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) VcpuCountInput() *TfCapacityProvider_VcpuCountProperty {
	var returns *TfCapacityProvider_VcpuCountProperty
	_jsii_.Get(
		j,
		"vcpuCountInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCapacityProvider_InstanceRequirementsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCapacityProvider_InstanceRequirementsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCapacityProvider_InstanceRequirementsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.InstanceRequirementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCapacityProvider_InstanceRequirementsPropertyOutputReference_Override(t TfCapacityProvider_InstanceRequirementsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.InstanceRequirementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetAllowedInstanceTypes(val *[]*string) {
	if err := j.validateSetAllowedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetInternalValue(val *TfCapacityProvider_InstanceRequirementsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64) {
	if err := j.validateSetMaxSpotPriceAsPercentageOfOptimalOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutAcceleratorCount(value *TfCapacityProvider_AcceleratorCountProperty) {
	if err := t.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutAcceleratorTotalMemoryMib(value *TfCapacityProvider_AcceleratorTotalMemoryMibProperty) {
	if err := t.validatePutAcceleratorTotalMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAcceleratorTotalMemoryMib",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutBaselineEbsBandwidthMbps(value *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty) {
	if err := t.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutMemoryGibPerVcpu(value *TfCapacityProvider_MemoryGibPerVcpuProperty) {
	if err := t.validatePutMemoryGibPerVcpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMemoryGibPerVcpu",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutMemoryMib(value *TfCapacityProvider_MemoryMibProperty) {
	if err := t.validatePutMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMemoryMib",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutNetworkBandwidthGbps(value *TfCapacityProvider_NetworkBandwidthGbpsProperty) {
	if err := t.validatePutNetworkBandwidthGbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkBandwidthGbps",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutNetworkInterfaceCount(value *TfCapacityProvider_NetworkInterfaceCountProperty) {
	if err := t.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutTotalLocalStorageGb(value *TfCapacityProvider_TotalLocalStorageGbProperty) {
	if err := t.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) PutVcpuCount(value *TfCapacityProvider_VcpuCountProperty) {
	if err := t.validatePutVcpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVcpuCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAcceleratorTotalMemoryMib() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorTotalMemoryMib",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetAllowedInstanceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedInstanceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		t,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		t,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		t,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		t,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		t,
		"resetMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetNetworkBandwidthGbps() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkBandwidthGbps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		t,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceRequirementsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

