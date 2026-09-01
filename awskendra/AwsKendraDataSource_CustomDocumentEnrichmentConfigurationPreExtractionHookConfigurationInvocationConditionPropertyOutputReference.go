package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference interface {
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
	ConditionOnValue() AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference
	// Experimental.
	ConditionOnValueInput() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty
	// Experimental.
	SetInternalValue(val *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty)
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
	PutConditionOnValue(value *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty)
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

// The jsii proxy struct for AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference
type jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionDocumentAttributeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionDocumentAttributeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionDocumentAttributeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionOnValue() AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference {
	var returns AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference
	_jsii_.Get(
		j,
		"conditionOnValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ConditionOnValueInput() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty {
	var returns *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	_jsii_.Get(
		j,
		"conditionOnValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InternalValue() *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty {
	var returns *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference_Override(a AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetConditionDocumentAttributeKey(val *string) {
	if err := j.validateSetConditionDocumentAttributeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionDocumentAttributeKey",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetInternalValue(val *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) PutConditionOnValue(value *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty) {
	if err := a.validatePutConditionOnValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConditionOnValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ResetConditionOnValue() {
	_jsii_.InvokeVoid(
		a,
		"resetConditionOnValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

