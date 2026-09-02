package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AsnRanges() *[]*string
	// Experimental.
	SetAsnRanges(val *[]*string)
	// Experimental.
	AsnRangesInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DnsSupport() interface{}
	// Experimental.
	SetDnsSupport(val interface{})
	// Experimental.
	DnsSupportInput() interface{}
	// Experimental.
	EdgeLocations() DataTfCoreNetworkPolicyDocument_EdgeLocationsPropertyList
	// Experimental.
	EdgeLocationsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InsideCidrBlocks() *[]*string
	// Experimental.
	SetInsideCidrBlocks(val *[]*string)
	// Experimental.
	InsideCidrBlocksInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SecurityGroupReferencingSupport() interface{}
	// Experimental.
	SetSecurityGroupReferencingSupport(val interface{})
	// Experimental.
	SecurityGroupReferencingSupportInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpnEcmpSupport() interface{}
	// Experimental.
	SetVpnEcmpSupport(val interface{})
	// Experimental.
	VpnEcmpSupportInput() interface{}
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
	PutEdgeLocations(value interface{})
	// Experimental.
	ResetDnsSupport()
	// Experimental.
	ResetInsideCidrBlocks()
	// Experimental.
	ResetSecurityGroupReferencingSupport()
	// Experimental.
	ResetVpnEcmpSupport()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference
type jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) AsnRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asnRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) AsnRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asnRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) DnsSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) DnsSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) EdgeLocations() DataTfCoreNetworkPolicyDocument_EdgeLocationsPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_EdgeLocationsPropertyList
	_jsii_.Get(
		j,
		"edgeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) EdgeLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) InsideCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insideCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) InsideCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insideCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) SecurityGroupReferencingSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityGroupReferencingSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) SecurityGroupReferencingSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityGroupReferencingSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) VpnEcmpSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpnEcmpSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) VpnEcmpSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpnEcmpSupportInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.CoreNetworkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference_Override(d DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.CoreNetworkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetAsnRanges(val *[]*string) {
	if err := j.validateSetAsnRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asnRanges",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetDnsSupport(val interface{}) {
	if err := j.validateSetDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSupport",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetInsideCidrBlocks(val *[]*string) {
	if err := j.validateSetInsideCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insideCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetSecurityGroupReferencingSupport(val interface{}) {
	if err := j.validateSetSecurityGroupReferencingSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupReferencingSupport",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference)SetVpnEcmpSupport(val interface{}) {
	if err := j.validateSetVpnEcmpSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnEcmpSupport",
		val,
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) PutEdgeLocations(value interface{}) {
	if err := d.validatePutEdgeLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEdgeLocations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ResetDnsSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetDnsSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ResetInsideCidrBlocks() {
	_jsii_.InvokeVoid(
		d,
		"resetInsideCidrBlocks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ResetSecurityGroupReferencingSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroupReferencingSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ResetVpnEcmpSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetVpnEcmpSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

