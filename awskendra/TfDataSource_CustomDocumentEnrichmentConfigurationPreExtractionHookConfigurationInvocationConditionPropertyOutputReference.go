package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference interface {
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
	// Experimental.
	ConditionDocumentAttributeKey() *string
	// Experimental.
	SetConditionDocumentAttributeKey(val *string)
	// Experimental.
	ConditionDocumentAttributeKeyInput() *string
	// Experimental.
	ConditionOnValue() TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference
	// Experimental.
	ConditionOnValueInput() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty)
	// Experimental.
	Operator() *string
	// Experimental.
	SetOperator(val *string)
	// Experimental.
	OperatorInput() *string
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
	PutConditionOnValue(value *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty)
	// Experimental.
	ResetConditionOnValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference
type jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionDocumentAttributeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionDocumentAttributeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionOnValue() TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference {
	var returns TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference
	_jsii_.Get(
		j,
		"conditionOnValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionOnValueInput() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty {
	var returns *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	_jsii_.Get(
		j,
		"conditionOnValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InternalValue() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty {
	var returns *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference_Override(t TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetConditionDocumentAttributeKey(val *string) {
	if err := j.validateSetConditionDocumentAttributeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionDocumentAttributeKey",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetInternalValue(val *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) PutConditionOnValue(value *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty) {
	if err := t.validatePutConditionOnValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConditionOnValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ResetConditionOnValue() {
	_jsii_.InvokeVoid(
		t,
		"resetConditionOnValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

