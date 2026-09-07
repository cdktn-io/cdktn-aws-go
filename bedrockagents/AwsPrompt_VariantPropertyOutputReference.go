package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPrompt_VariantPropertyOutputReference interface {
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
	GenAiResource() AwsPrompt_GenAiResourcePropertyList
	// Experimental.
	GenAiResourceInput() interface{}
	// Experimental.
	InferenceConfiguration() AwsPrompt_InferenceConfigurationPropertyList
	// Experimental.
	InferenceConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Metadata() AwsPrompt_MetadataPropertyList
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
	TemplateConfiguration() AwsPrompt_TemplateConfigurationPropertyList
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

// The jsii proxy struct for AwsPrompt_VariantPropertyOutputReference
type jsiiProxy_AwsPrompt_VariantPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) AdditionalModelRequestFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) AdditionalModelRequestFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GenAiResource() AwsPrompt_GenAiResourcePropertyList {
	var returns AwsPrompt_GenAiResourcePropertyList
	_jsii_.Get(
		j,
		"genAiResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GenAiResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genAiResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) InferenceConfiguration() AwsPrompt_InferenceConfigurationPropertyList {
	var returns AwsPrompt_InferenceConfigurationPropertyList
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) Metadata() AwsPrompt_MetadataPropertyList {
	var returns AwsPrompt_MetadataPropertyList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TemplateConfiguration() AwsPrompt_TemplateConfigurationPropertyList {
	var returns AwsPrompt_TemplateConfigurationPropertyList
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPrompt_VariantPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPrompt_VariantPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPrompt_VariantPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPrompt_VariantPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsPrompt.VariantPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPrompt_VariantPropertyOutputReference_Override(a AwsPrompt_VariantPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsPrompt.VariantPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetAdditionalModelRequestFields(val *string) {
	if err := j.validateSetAdditionalModelRequestFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalModelRequestFields",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetTemplateType(val *string) {
	if err := j.validateSetTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateType",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPrompt_VariantPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) PutGenAiResource(value interface{}) {
	if err := a.validatePutGenAiResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGenAiResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) PutInferenceConfiguration(value interface{}) {
	if err := a.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) PutMetadata(value interface{}) {
	if err := a.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) PutTemplateConfiguration(value interface{}) {
	if err := a.validatePutTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTemplateConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetAdditionalModelRequestFields() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalModelRequestFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetGenAiResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGenAiResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		a,
		"resetModelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ResetTemplateConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPrompt_VariantPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

