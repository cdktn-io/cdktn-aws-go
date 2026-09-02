package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNodeGroup_UpdateConfigPropertyOutputReference interface {
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
	InternalValue() *TfNodeGroup_UpdateConfigProperty
	// Experimental.
	SetInternalValue(val *TfNodeGroup_UpdateConfigProperty)
	// Experimental.
	MaxUnavailable() *float64
	// Experimental.
	SetMaxUnavailable(val *float64)
	// Experimental.
	MaxUnavailableInput() *float64
	// Experimental.
	MaxUnavailablePercentage() *float64
	// Experimental.
	SetMaxUnavailablePercentage(val *float64)
	// Experimental.
	MaxUnavailablePercentageInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UpdateStrategy() *string
	// Experimental.
	SetUpdateStrategy(val *string)
	// Experimental.
	UpdateStrategyInput() *string
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
	ResetMaxUnavailable()
	// Experimental.
	ResetMaxUnavailablePercentage()
	// Experimental.
	ResetUpdateStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfNodeGroup_UpdateConfigPropertyOutputReference
type jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) InternalValue() *TfNodeGroup_UpdateConfigProperty {
	var returns *TfNodeGroup_UpdateConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) MaxUnavailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) MaxUnavailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) MaxUnavailablePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) MaxUnavailablePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) UpdateStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) UpdateStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateStrategyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfNodeGroup_UpdateConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfNodeGroup_UpdateConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfNodeGroup_UpdateConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.UpdateConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfNodeGroup_UpdateConfigPropertyOutputReference_Override(t TfNodeGroup_UpdateConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.UpdateConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetInternalValue(val *TfNodeGroup_UpdateConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetMaxUnavailable(val *float64) {
	if err := j.validateSetMaxUnavailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailable",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetMaxUnavailablePercentage(val *float64) {
	if err := j.validateSetMaxUnavailablePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailablePercentage",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference)SetUpdateStrategy(val *string) {
	if err := j.validateSetUpdateStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateStrategy",
		val,
	)
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ResetMaxUnavailable() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxUnavailable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ResetMaxUnavailablePercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxUnavailablePercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ResetUpdateStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdateStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfNodeGroup_UpdateConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

