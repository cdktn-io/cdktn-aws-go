package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Analyzer() *string
	// Experimental.
	SetAnalyzer(val *string)
	// Experimental.
	AnalyzerInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ResultField() *string
	// Experimental.
	SetResultField(val *string)
	// Experimental.
	ResultFieldInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference
type jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) Analyzer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) AnalyzerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ResultField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ResultFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfTrafficPolicy.PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference_Override(t TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfTrafficPolicy.PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetAnalyzer(val *string) {
	if err := j.validateSetAnalyzerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyzer",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetResultField(val *string) {
	if err := j.validateSetResultFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultField",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateAnalysisPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

