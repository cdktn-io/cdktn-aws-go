package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference interface {
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
	MacAddress() *string
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
	SecondaryInterfaceId() *string
	// Experimental.
	SecondaryNetworkId() *string
	// Experimental.
	SecondarySubnetId() *string
	// Experimental.
	SetSecondarySubnetId(val *string)
	// Experimental.
	SecondarySubnetIdInput() *string
	// Experimental.
	SourceDestCheck() cdktn.IResolvable
	// Experimental.
	Status() *string
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
	ResetPrivateIpAddressCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference
type jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) PrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) PrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) SecondaryInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) SecondaryNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) SecondarySubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) SecondarySubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) SourceDestCheck() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfSpotInstanceRequest.SecondaryNetworkInterfacePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference_Override(t TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfSpotInstanceRequest.SecondaryNetworkInterfacePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetPrivateIpAddressCount(val *float64) {
	if err := j.validateSetPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetSecondarySubnetId(val *string) {
	if err := j.validateSetSecondarySubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySubnetId",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ResetPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSpotInstanceRequest_SecondaryNetworkInterfacePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

