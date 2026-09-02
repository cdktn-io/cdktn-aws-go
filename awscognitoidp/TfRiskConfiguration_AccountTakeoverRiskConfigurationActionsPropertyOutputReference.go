package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	HighAction() TfRiskConfiguration_HighActionPropertyOutputReference
	// Experimental.
	HighActionInput() *TfRiskConfiguration_HighActionProperty
	// Experimental.
	InternalValue() *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	// Experimental.
	SetInternalValue(val *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty)
	// Experimental.
	LowAction() TfRiskConfiguration_LowActionPropertyOutputReference
	// Experimental.
	LowActionInput() *TfRiskConfiguration_LowActionProperty
	// Experimental.
	MediumAction() TfRiskConfiguration_MediumActionPropertyOutputReference
	// Experimental.
	MediumActionInput() *TfRiskConfiguration_MediumActionProperty
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
	PutHighAction(value *TfRiskConfiguration_HighActionProperty)
	// Experimental.
	PutLowAction(value *TfRiskConfiguration_LowActionProperty)
	// Experimental.
	PutMediumAction(value *TfRiskConfiguration_MediumActionProperty)
	// Experimental.
	ResetHighAction()
	// Experimental.
	ResetLowAction()
	// Experimental.
	ResetMediumAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference
type jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) HighAction() TfRiskConfiguration_HighActionPropertyOutputReference {
	var returns TfRiskConfiguration_HighActionPropertyOutputReference
	_jsii_.Get(
		j,
		"highAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) HighActionInput() *TfRiskConfiguration_HighActionProperty {
	var returns *TfRiskConfiguration_HighActionProperty
	_jsii_.Get(
		j,
		"highActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InternalValue() *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty {
	var returns *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) LowAction() TfRiskConfiguration_LowActionPropertyOutputReference {
	var returns TfRiskConfiguration_LowActionPropertyOutputReference
	_jsii_.Get(
		j,
		"lowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) LowActionInput() *TfRiskConfiguration_LowActionProperty {
	var returns *TfRiskConfiguration_LowActionProperty
	_jsii_.Get(
		j,
		"lowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) MediumAction() TfRiskConfiguration_MediumActionPropertyOutputReference {
	var returns TfRiskConfiguration_MediumActionPropertyOutputReference
	_jsii_.Get(
		j,
		"mediumAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) MediumActionInput() *TfRiskConfiguration_MediumActionProperty {
	var returns *TfRiskConfiguration_MediumActionProperty
	_jsii_.Get(
		j,
		"mediumActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.AccountTakeoverRiskConfigurationActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference_Override(t TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.AccountTakeoverRiskConfigurationActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetInternalValue(val *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutHighAction(value *TfRiskConfiguration_HighActionProperty) {
	if err := t.validatePutHighActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHighAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutLowAction(value *TfRiskConfiguration_LowActionProperty) {
	if err := t.validatePutLowActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLowAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) PutMediumAction(value *TfRiskConfiguration_MediumActionProperty) {
	if err := t.validatePutMediumActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMediumAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetHighAction() {
	_jsii_.InvokeVoid(
		t,
		"resetHighAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetLowAction() {
	_jsii_.InvokeVoid(
		t,
		"resetLowAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ResetMediumAction() {
	_jsii_.InvokeVoid(
		t,
		"resetMediumAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

