package ec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsImageRecipe_BlockDeviceMappingPropertyOutputReference interface {
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
	DeviceName() *string
	// Experimental.
	SetDeviceName(val *string)
	// Experimental.
	DeviceNameInput() *string
	// Experimental.
	Ebs() AwsImageRecipe_EbsPropertyOutputReference
	// Experimental.
	EbsInput() *AwsImageRecipe_EbsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NoDevice() interface{}
	// Experimental.
	SetNoDevice(val interface{})
	// Experimental.
	NoDeviceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VirtualName() *string
	// Experimental.
	SetVirtualName(val *string)
	// Experimental.
	VirtualNameInput() *string
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
	PutEbs(value *AwsImageRecipe_EbsProperty)
	// Experimental.
	ResetDeviceName()
	// Experimental.
	ResetEbs()
	// Experimental.
	ResetNoDevice()
	// Experimental.
	ResetVirtualName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsImageRecipe_BlockDeviceMappingPropertyOutputReference
type jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) Ebs() AwsImageRecipe_EbsPropertyOutputReference {
	var returns AwsImageRecipe_EbsPropertyOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) EbsInput() *AwsImageRecipe_EbsProperty {
	var returns *AwsImageRecipe_EbsProperty
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) NoDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) NoDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) VirtualNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsImageRecipe_BlockDeviceMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsImageRecipe_BlockDeviceMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsImageRecipe_BlockDeviceMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImageRecipe.BlockDeviceMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsImageRecipe_BlockDeviceMappingPropertyOutputReference_Override(a AwsImageRecipe_BlockDeviceMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImageRecipe.BlockDeviceMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetNoDevice(val interface{}) {
	if err := j.validateSetNoDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDevice",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference)SetVirtualName(val *string) {
	if err := j.validateSetVirtualNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualName",
		val,
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) PutEbs(value *AwsImageRecipe_EbsProperty) {
	if err := a.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ResetEbs() {
	_jsii_.InvokeVoid(
		a,
		"resetEbs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ResetNoDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetNoDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ResetVirtualName() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsImageRecipe_BlockDeviceMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

