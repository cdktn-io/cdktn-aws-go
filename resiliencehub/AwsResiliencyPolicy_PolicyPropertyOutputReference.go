package resiliencehub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/resiliencehub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/resiliencehub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResiliencyPolicy_PolicyPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Az() AwsResiliencyPolicy_AzPropertyList
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
	Hardware() AwsResiliencyPolicy_HardwarePropertyList
	// Experimental.
	HardwareInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Region() AwsResiliencyPolicy_RegionPropertyList
	// Experimental.
	RegionInput() interface{}
	// Experimental.
	SoftwareAttribute() AwsResiliencyPolicy_SoftwarePropertyList
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

// The jsii proxy struct for AwsResiliencyPolicy_PolicyPropertyOutputReference
type jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) Az() AwsResiliencyPolicy_AzPropertyList {
	var returns AwsResiliencyPolicy_AzPropertyList
	_jsii_.Get(
		j,
		"az",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) AzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) Hardware() AwsResiliencyPolicy_HardwarePropertyList {
	var returns AwsResiliencyPolicy_HardwarePropertyList
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) HardwareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) Region() AwsResiliencyPolicy_RegionPropertyList {
	var returns AwsResiliencyPolicy_RegionPropertyList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) RegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) SoftwareAttribute() AwsResiliencyPolicy_SoftwarePropertyList {
	var returns AwsResiliencyPolicy_SoftwarePropertyList
	_jsii_.Get(
		j,
		"softwareAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) SoftwareAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softwareAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResiliencyPolicy_PolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsResiliencyPolicy_PolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResiliencyPolicy_PolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-resilience-hub.AwsResiliencyPolicy.PolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResiliencyPolicy_PolicyPropertyOutputReference_Override(a AwsResiliencyPolicy_PolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-resilience-hub.AwsResiliencyPolicy.PolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) PutAz(value interface{}) {
	if err := a.validatePutAzParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAz",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) PutHardware(value interface{}) {
	if err := a.validatePutHardwareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHardware",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) PutRegion(value interface{}) {
	if err := a.validatePutRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) PutSoftwareAttribute(value interface{}) {
	if err := a.validatePutSoftwareAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoftwareAttribute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ResetAz() {
	_jsii_.InvokeVoid(
		a,
		"resetAz",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ResetHardware() {
	_jsii_.InvokeVoid(
		a,
		"resetHardware",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ResetSoftwareAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetSoftwareAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

