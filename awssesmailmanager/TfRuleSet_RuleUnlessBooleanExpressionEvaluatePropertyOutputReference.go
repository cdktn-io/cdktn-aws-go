package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Analysis() TfRuleSet_RuleUnlessBooleanExpressionEvaluateAnalysisPropertyList
	// Experimental.
	AnalysisInput() interface{}
	// Experimental.
	Attribute() *string
	// Experimental.
	SetAttribute(val *string)
	// Experimental.
	AttributeInput() *string
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
	IsInAddressList() TfRuleSet_RuleUnlessBooleanExpressionEvaluateIsInAddressListPropertyList
	// Experimental.
	IsInAddressListInput() interface{}
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
	PutAnalysis(value interface{})
	// Experimental.
	PutIsInAddressList(value interface{})
	// Experimental.
	ResetAnalysis()
	// Experimental.
	ResetAttribute()
	// Experimental.
	ResetIsInAddressList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference
type jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) Analysis() TfRuleSet_RuleUnlessBooleanExpressionEvaluateAnalysisPropertyList {
	var returns TfRuleSet_RuleUnlessBooleanExpressionEvaluateAnalysisPropertyList
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) AnalysisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) IsInAddressList() TfRuleSet_RuleUnlessBooleanExpressionEvaluateIsInAddressListPropertyList {
	var returns TfRuleSet_RuleUnlessBooleanExpressionEvaluateIsInAddressListPropertyList
	_jsii_.Get(
		j,
		"isInAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) IsInAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.RuleUnlessBooleanExpressionEvaluatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference_Override(t TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.RuleUnlessBooleanExpressionEvaluatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) PutAnalysis(value interface{}) {
	if err := t.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnalysis",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) PutIsInAddressList(value interface{}) {
	if err := t.validatePutIsInAddressListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIsInAddressList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		t,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ResetIsInAddressList() {
	_jsii_.InvokeVoid(
		t,
		"resetIsInAddressList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleSet_RuleUnlessBooleanExpressionEvaluatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

