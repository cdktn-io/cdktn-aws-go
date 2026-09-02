package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BedrockKnowledgeBaseArn() *string
	// Experimental.
	SetBedrockKnowledgeBaseArn(val *string)
	// Experimental.
	BedrockKnowledgeBaseArnInput() *string
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
	ExactResponse() interface{}
	// Experimental.
	SetExactResponse(val interface{})
	// Experimental.
	ExactResponseFields() TfIntent_QnaIntentConfigurationDataSourceConfigurationBedrockKnowledgeStoreConfigurationExactResponseFieldsPropertyList
	// Experimental.
	ExactResponseFieldsInput() interface{}
	// Experimental.
	ExactResponseInput() interface{}
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
	PutExactResponseFields(value interface{})
	// Experimental.
	ResetExactResponse()
	// Experimental.
	ResetExactResponseFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference
type jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) BedrockKnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockKnowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) BedrockKnowledgeBaseArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockKnowledgeBaseArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ExactResponse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ExactResponseFields() TfIntent_QnaIntentConfigurationDataSourceConfigurationBedrockKnowledgeStoreConfigurationExactResponseFieldsPropertyList {
	var returns TfIntent_QnaIntentConfigurationDataSourceConfigurationBedrockKnowledgeStoreConfigurationExactResponseFieldsPropertyList
	_jsii_.Get(
		j,
		"exactResponseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ExactResponseFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponseFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ExactResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.BedrockKnowledgeStoreConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference_Override(t TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.BedrockKnowledgeStoreConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetBedrockKnowledgeBaseArn(val *string) {
	if err := j.validateSetBedrockKnowledgeBaseArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bedrockKnowledgeBaseArn",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetExactResponse(val interface{}) {
	if err := j.validateSetExactResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactResponse",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) PutExactResponseFields(value interface{}) {
	if err := t.validatePutExactResponseFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExactResponseFields",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ResetExactResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetExactResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ResetExactResponseFields() {
	_jsii_.InvokeVoid(
		t,
		"resetExactResponseFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_BedrockKnowledgeStoreConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

