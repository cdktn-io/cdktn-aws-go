package awsiamaccessanalyzer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiamaccessanalyzer/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiamaccessanalyzer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAnalyzer_UnusedAccessPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AnalysisRule() TfAnalyzer_ConfigurationUnusedAccessAnalysisRulePropertyOutputReference
	// Experimental.
	AnalysisRuleInput() *TfAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty
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
	InternalValue() *TfAnalyzer_UnusedAccessProperty
	// Experimental.
	SetInternalValue(val *TfAnalyzer_UnusedAccessProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnusedAccessAge() *float64
	// Experimental.
	SetUnusedAccessAge(val *float64)
	// Experimental.
	UnusedAccessAgeInput() *float64
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
	PutAnalysisRule(value *TfAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty)
	// Experimental.
	ResetAnalysisRule()
	// Experimental.
	ResetUnusedAccessAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAnalyzer_UnusedAccessPropertyOutputReference
type jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) AnalysisRule() TfAnalyzer_ConfigurationUnusedAccessAnalysisRulePropertyOutputReference {
	var returns TfAnalyzer_ConfigurationUnusedAccessAnalysisRulePropertyOutputReference
	_jsii_.Get(
		j,
		"analysisRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) AnalysisRuleInput() *TfAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty {
	var returns *TfAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty
	_jsii_.Get(
		j,
		"analysisRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) InternalValue() *TfAnalyzer_UnusedAccessProperty {
	var returns *TfAnalyzer_UnusedAccessProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) UnusedAccessAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unusedAccessAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) UnusedAccessAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unusedAccessAgeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAnalyzer_UnusedAccessPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAnalyzer_UnusedAccessPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAnalyzer_UnusedAccessPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iam-access-analyzer.TfAnalyzer.UnusedAccessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAnalyzer_UnusedAccessPropertyOutputReference_Override(t TfAnalyzer_UnusedAccessPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iam-access-analyzer.TfAnalyzer.UnusedAccessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetInternalValue(val *TfAnalyzer_UnusedAccessProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference)SetUnusedAccessAge(val *float64) {
	if err := j.validateSetUnusedAccessAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unusedAccessAge",
		val,
	)
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) PutAnalysisRule(value *TfAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty) {
	if err := t.validatePutAnalysisRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnalysisRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ResetAnalysisRule() {
	_jsii_.InvokeVoid(
		t,
		"resetAnalysisRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ResetUnusedAccessAge() {
	_jsii_.InvokeVoid(
		t,
		"resetUnusedAccessAge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAnalyzer_UnusedAccessPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

