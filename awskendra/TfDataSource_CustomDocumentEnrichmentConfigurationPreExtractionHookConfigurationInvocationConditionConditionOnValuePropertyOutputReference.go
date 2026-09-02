package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference interface {
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
	DateValue() *string
	// Experimental.
	SetDateValue(val *string)
	// Experimental.
	DateValueInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty)
	// Experimental.
	LongValue() *float64
	// Experimental.
	SetLongValue(val *float64)
	// Experimental.
	LongValueInput() *float64
	// Experimental.
	StringListValue() *[]*string
	// Experimental.
	SetStringListValue(val *[]*string)
	// Experimental.
	StringListValueInput() *[]*string
	// Experimental.
	StringValue() *string
	// Experimental.
	SetStringValue(val *string)
	// Experimental.
	StringValueInput() *string
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
	ResetDateValue()
	// Experimental.
	ResetLongValue()
	// Experimental.
	ResetStringListValue()
	// Experimental.
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference
type jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) DateValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) DateValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) InternalValue() *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty {
	var returns *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) LongValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) LongValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) StringListValueInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference_Override(t TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetDateValue(val *string) {
	if err := j.validateSetDateValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetInternalValue(val *TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValueProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetLongValue(val *float64) {
	if err := j.validateSetLongValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetStringListValue(val *[]*string) {
	if err := j.validateSetStringListValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringListValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ResetDateValue() {
	_jsii_.InvokeVoid(
		t,
		"resetDateValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ResetLongValue() {
	_jsii_.InvokeVoid(
		t,
		"resetLongValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		t,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		t,
		"resetStringValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_CustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValuePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

