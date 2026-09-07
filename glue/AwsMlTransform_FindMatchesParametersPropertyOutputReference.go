package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMlTransform_FindMatchesParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccuracyCostTradeOff() *float64
	// Experimental.
	SetAccuracyCostTradeOff(val *float64)
	// Experimental.
	AccuracyCostTradeOffInput() *float64
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
	EnforceProvidedLabels() interface{}
	// Experimental.
	SetEnforceProvidedLabels(val interface{})
	// Experimental.
	EnforceProvidedLabelsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMlTransform_FindMatchesParametersProperty
	// Experimental.
	SetInternalValue(val *AwsMlTransform_FindMatchesParametersProperty)
	// Experimental.
	PrecisionRecallTradeOff() *float64
	// Experimental.
	SetPrecisionRecallTradeOff(val *float64)
	// Experimental.
	PrecisionRecallTradeOffInput() *float64
	// Experimental.
	PrimaryKeyColumnName() *string
	// Experimental.
	SetPrimaryKeyColumnName(val *string)
	// Experimental.
	PrimaryKeyColumnNameInput() *string
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
	ResetAccuracyCostTradeOff()
	// Experimental.
	ResetEnforceProvidedLabels()
	// Experimental.
	ResetPrecisionRecallTradeOff()
	// Experimental.
	ResetPrimaryKeyColumnName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMlTransform_FindMatchesParametersPropertyOutputReference
type jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) AccuracyCostTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) AccuracyCostTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) EnforceProvidedLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) EnforceProvidedLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) InternalValue() *AwsMlTransform_FindMatchesParametersProperty {
	var returns *AwsMlTransform_FindMatchesParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) PrecisionRecallTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) PrecisionRecallTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) PrimaryKeyColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) PrimaryKeyColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMlTransform_FindMatchesParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMlTransform_FindMatchesParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMlTransform_FindMatchesParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsMlTransform.FindMatchesParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMlTransform_FindMatchesParametersPropertyOutputReference_Override(a AwsMlTransform_FindMatchesParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsMlTransform.FindMatchesParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetAccuracyCostTradeOff(val *float64) {
	if err := j.validateSetAccuracyCostTradeOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accuracyCostTradeOff",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetEnforceProvidedLabels(val interface{}) {
	if err := j.validateSetEnforceProvidedLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceProvidedLabels",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetInternalValue(val *AwsMlTransform_FindMatchesParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetPrecisionRecallTradeOff(val *float64) {
	if err := j.validateSetPrecisionRecallTradeOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precisionRecallTradeOff",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetPrimaryKeyColumnName(val *string) {
	if err := j.validateSetPrimaryKeyColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumnName",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ResetAccuracyCostTradeOff() {
	_jsii_.InvokeVoid(
		a,
		"resetAccuracyCostTradeOff",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ResetEnforceProvidedLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceProvidedLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ResetPrecisionRecallTradeOff() {
	_jsii_.InvokeVoid(
		a,
		"resetPrecisionRecallTradeOff",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ResetPrimaryKeyColumnName() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryKeyColumnName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMlTransform_FindMatchesParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

