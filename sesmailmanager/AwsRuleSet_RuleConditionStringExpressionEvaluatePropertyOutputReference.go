package sesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Analysis() AwsRuleSet_RuleConditionStringExpressionEvaluateAnalysisPropertyList
	// Experimental.
	AnalysisInput() interface{}
	// Experimental.
	Attribute() *string
	// Experimental.
	SetAttribute(val *string)
	// Experimental.
	AttributeInput() *string
	// Experimental.
	ClientCertificateAttribute() *string
	// Experimental.
	SetClientCertificateAttribute(val *string)
	// Experimental.
	ClientCertificateAttributeInput() *string
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
	MimeHeaderAttribute() *string
	// Experimental.
	SetMimeHeaderAttribute(val *string)
	// Experimental.
	MimeHeaderAttributeInput() *string
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
	ResetAnalysis()
	// Experimental.
	ResetAttribute()
	// Experimental.
	ResetClientCertificateAttribute()
	// Experimental.
	ResetMimeHeaderAttribute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference
type jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) Analysis() AwsRuleSet_RuleConditionStringExpressionEvaluateAnalysisPropertyList {
	var returns AwsRuleSet_RuleConditionStringExpressionEvaluateAnalysisPropertyList
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) AnalysisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ClientCertificateAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ClientCertificateAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) MimeHeaderAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeHeaderAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) MimeHeaderAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeHeaderAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsRuleSet.RuleConditionStringExpressionEvaluatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference_Override(a AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsRuleSet.RuleConditionStringExpressionEvaluatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetClientCertificateAttribute(val *string) {
	if err := j.validateSetClientCertificateAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetMimeHeaderAttribute(val *string) {
	if err := j.validateSetMimeHeaderAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mimeHeaderAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) PutAnalysis(value interface{}) {
	if err := a.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnalysis",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ResetClientCertificateAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ResetMimeHeaderAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetMimeHeaderAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRuleSet_RuleConditionStringExpressionEvaluatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

