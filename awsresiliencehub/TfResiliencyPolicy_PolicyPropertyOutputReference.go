package awsresiliencehub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsresiliencehub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsresiliencehub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResiliencyPolicy_PolicyPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Az() TfResiliencyPolicy_AzPropertyList
	// Experimental.
	AzInput() interface{}
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
	Fqn() *string
	// Experimental.
	Hardware() TfResiliencyPolicy_HardwarePropertyList
	// Experimental.
	HardwareInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Region() TfResiliencyPolicy_RegionPropertyList
	// Experimental.
	RegionInput() interface{}
	// Experimental.
	SoftwareAttribute() TfResiliencyPolicy_SoftwarePropertyList
	// Experimental.
	SoftwareAttributeInput() interface{}
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
	PutAz(value interface{})
	// Experimental.
	PutHardware(value interface{})
	// Experimental.
	PutRegion(value interface{})
	// Experimental.
	PutSoftwareAttribute(value interface{})
	// Experimental.
	ResetAz()
	// Experimental.
	ResetHardware()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSoftwareAttribute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResiliencyPolicy_PolicyPropertyOutputReference
type jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) Az() TfResiliencyPolicy_AzPropertyList {
	var returns TfResiliencyPolicy_AzPropertyList
	_jsii_.Get(
		j,
		"az",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) AzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) Hardware() TfResiliencyPolicy_HardwarePropertyList {
	var returns TfResiliencyPolicy_HardwarePropertyList
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) HardwareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) Region() TfResiliencyPolicy_RegionPropertyList {
	var returns TfResiliencyPolicy_RegionPropertyList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) RegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) SoftwareAttribute() TfResiliencyPolicy_SoftwarePropertyList {
	var returns TfResiliencyPolicy_SoftwarePropertyList
	_jsii_.Get(
		j,
		"softwareAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) SoftwareAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softwareAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResiliencyPolicy_PolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfResiliencyPolicy_PolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResiliencyPolicy_PolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-resilience-hub.TfResiliencyPolicy.PolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResiliencyPolicy_PolicyPropertyOutputReference_Override(t TfResiliencyPolicy_PolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-resilience-hub.TfResiliencyPolicy.PolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) PutAz(value interface{}) {
	if err := t.validatePutAzParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAz",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) PutHardware(value interface{}) {
	if err := t.validatePutHardwareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHardware",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) PutRegion(value interface{}) {
	if err := t.validatePutRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) PutSoftwareAttribute(value interface{}) {
	if err := t.validatePutSoftwareAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSoftwareAttribute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ResetAz() {
	_jsii_.InvokeVoid(
		t,
		"resetAz",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ResetHardware() {
	_jsii_.InvokeVoid(
		t,
		"resetHardware",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ResetSoftwareAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetSoftwareAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResiliencyPolicy_PolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

