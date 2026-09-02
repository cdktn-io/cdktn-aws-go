package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPrompt_VariantPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalModelRequestFields() *string
	// Experimental.
	SetAdditionalModelRequestFields(val *string)
	// Experimental.
	AdditionalModelRequestFieldsInput() *string
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
	GenAiResource() TfPrompt_GenAiResourcePropertyList
	// Experimental.
	GenAiResourceInput() interface{}
	// Experimental.
	InferenceConfiguration() TfPrompt_InferenceConfigurationPropertyList
	// Experimental.
	InferenceConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Metadata() TfPrompt_MetadataPropertyList
	// Experimental.
	MetadataInput() interface{}
	// Experimental.
	ModelId() *string
	// Experimental.
	SetModelId(val *string)
	// Experimental.
	ModelIdInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	TemplateConfiguration() TfPrompt_TemplateConfigurationPropertyList
	// Experimental.
	TemplateConfigurationInput() interface{}
	// Experimental.
	TemplateType() *string
	// Experimental.
	SetTemplateType(val *string)
	// Experimental.
	TemplateTypeInput() *string
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
	PutGenAiResource(value interface{})
	// Experimental.
	PutInferenceConfiguration(value interface{})
	// Experimental.
	PutMetadata(value interface{})
	// Experimental.
	PutTemplateConfiguration(value interface{})
	// Experimental.
	ResetAdditionalModelRequestFields()
	// Experimental.
	ResetGenAiResource()
	// Experimental.
	ResetInferenceConfiguration()
	// Experimental.
	ResetMetadata()
	// Experimental.
	ResetModelId()
	// Experimental.
	ResetTemplateConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPrompt_VariantPropertyOutputReference
type jsiiProxy_TfPrompt_VariantPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) AdditionalModelRequestFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) AdditionalModelRequestFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GenAiResource() TfPrompt_GenAiResourcePropertyList {
	var returns TfPrompt_GenAiResourcePropertyList
	_jsii_.Get(
		j,
		"genAiResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GenAiResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genAiResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) InferenceConfiguration() TfPrompt_InferenceConfigurationPropertyList {
	var returns TfPrompt_InferenceConfigurationPropertyList
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) Metadata() TfPrompt_MetadataPropertyList {
	var returns TfPrompt_MetadataPropertyList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TemplateConfiguration() TfPrompt_TemplateConfigurationPropertyList {
	var returns TfPrompt_TemplateConfigurationPropertyList
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPrompt_VariantPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPrompt_VariantPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPrompt_VariantPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPrompt_VariantPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfPrompt.VariantPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPrompt_VariantPropertyOutputReference_Override(t TfPrompt_VariantPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfPrompt.VariantPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetAdditionalModelRequestFields(val *string) {
	if err := j.validateSetAdditionalModelRequestFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalModelRequestFields",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetTemplateType(val *string) {
	if err := j.validateSetTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateType",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPrompt_VariantPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) PutGenAiResource(value interface{}) {
	if err := t.validatePutGenAiResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGenAiResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) PutInferenceConfiguration(value interface{}) {
	if err := t.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) PutMetadata(value interface{}) {
	if err := t.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadata",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) PutTemplateConfiguration(value interface{}) {
	if err := t.validatePutTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTemplateConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetAdditionalModelRequestFields() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalModelRequestFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetGenAiResource() {
	_jsii_.InvokeVoid(
		t,
		"resetGenAiResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		t,
		"resetModelId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ResetTemplateConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTemplateConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPrompt_VariantPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

