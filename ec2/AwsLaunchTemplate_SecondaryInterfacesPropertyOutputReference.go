package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DeleteOnTermination() interface{}
	// Experimental.
	SetDeleteOnTermination(val interface{})
	// Experimental.
	DeleteOnTerminationInput() interface{}
	// Experimental.
	DeviceIndex() *float64
	// Experimental.
	SetDeviceIndex(val *float64)
	// Experimental.
	DeviceIndexInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InterfaceType() *string
	// Experimental.
	SetInterfaceType(val *string)
	// Experimental.
	InterfaceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NetworkCardIndex() *float64
	// Experimental.
	SetNetworkCardIndex(val *float64)
	// Experimental.
	NetworkCardIndexInput() *float64
	// Experimental.
	PrivateIpAddressCount() *float64
	// Experimental.
	SetPrivateIpAddressCount(val *float64)
	// Experimental.
	PrivateIpAddressCountInput() *float64
	// Experimental.
	PrivateIpAddresses() *[]*string
	// Experimental.
	SetPrivateIpAddresses(val *[]*string)
	// Experimental.
	PrivateIpAddressesInput() *[]*string
	// Experimental.
	SecondarySubnetId() *string
	// Experimental.
	SetSecondarySubnetId(val *string)
	// Experimental.
	SecondarySubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetDeleteOnTermination()
	// Experimental.
	ResetDeviceIndex()
	// Experimental.
	ResetInterfaceType()
	// Experimental.
	ResetNetworkCardIndex()
	// Experimental.
	ResetPrivateIpAddressCount()
	// Experimental.
	ResetPrivateIpAddresses()
	// Experimental.
	ResetSecondarySubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference
type jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) PrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) PrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) PrivateIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) SecondarySubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) SecondarySubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLaunchTemplate_SecondaryInterfacesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.SecondaryInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference_Override(a AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.SecondaryInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetPrivateIpAddressCount(val *float64) {
	if err := j.validateSetPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetPrivateIpAddresses(val *[]*string) {
	if err := j.validateSetPrivateIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddresses",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetSecondarySubnetId(val *string) {
	if err := j.validateSetSecondarySubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySubnetId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetNetworkCardIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkCardIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetPrivateIpAddresses() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIpAddresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ResetSecondarySubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondarySubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_SecondaryInterfacesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

