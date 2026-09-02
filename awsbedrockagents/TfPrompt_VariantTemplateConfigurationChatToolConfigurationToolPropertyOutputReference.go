package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CachePoint() TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolCachePointPropertyList
	// Experimental.
	CachePointInput() interface{}
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ToolSpec() TfPrompt_ToolSpecPropertyList
	// Experimental.
	ToolSpecInput() interface{}
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
	PutCachePoint(value interface{})
	// Experimental.
	PutToolSpec(value interface{})
	// Experimental.
	ResetCachePoint()
	// Experimental.
	ResetToolSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference
type jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) CachePoint() TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolCachePointPropertyList {
	var returns TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolCachePointPropertyList
	_jsii_.Get(
		j,
		"cachePoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) CachePointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachePointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ToolSpec() TfPrompt_ToolSpecPropertyList {
	var returns TfPrompt_ToolSpecPropertyList
	_jsii_.Get(
		j,
		"toolSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ToolSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolSpecInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfPrompt.VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference_Override(t TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfPrompt.VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) PutCachePoint(value interface{}) {
	if err := t.validatePutCachePointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCachePoint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) PutToolSpec(value interface{}) {
	if err := t.validatePutToolSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putToolSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ResetCachePoint() {
	_jsii_.InvokeVoid(
		t,
		"resetCachePoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ResetToolSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetToolSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

