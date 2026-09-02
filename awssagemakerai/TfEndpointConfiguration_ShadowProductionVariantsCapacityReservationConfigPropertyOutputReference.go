package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityReservationPreference() *string
	// Experimental.
	SetCapacityReservationPreference(val *string)
	// Experimental.
	CapacityReservationPreferenceInput() *string
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
	InternalValue() *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
	// Experimental.
	SetInternalValue(val *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty)
	// Experimental.
	MlReservationArn() *string
	// Experimental.
	SetMlReservationArn(val *string)
	// Experimental.
	MlReservationArnInput() *string
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
	ResetCapacityReservationPreference()
	// Experimental.
	ResetMlReservationArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference
type jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CapacityReservationPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CapacityReservationPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InternalValue() *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty {
	var returns *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) MlReservationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlReservationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) MlReservationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlReservationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference_Override(t TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetCapacityReservationPreference(val *string) {
	if err := j.validateSetCapacityReservationPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationPreference",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetInternalValue(val *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetMlReservationArn(val *string) {
	if err := j.validateSetMlReservationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlReservationArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ResetCapacityReservationPreference() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationPreference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ResetMlReservationArn() {
	_jsii_.InvokeVoid(
		t,
		"resetMlReservationArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

