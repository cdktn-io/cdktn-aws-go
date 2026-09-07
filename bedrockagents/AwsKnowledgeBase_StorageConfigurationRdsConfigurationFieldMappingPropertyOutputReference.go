package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference interface {
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
	CustomMetadataField() *string
	// Experimental.
	SetCustomMetadataField(val *string)
	// Experimental.
	CustomMetadataFieldInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MetadataField() *string
	// Experimental.
	SetMetadataField(val *string)
	// Experimental.
	MetadataFieldInput() *string
	// Experimental.
	PrimaryKeyField() *string
	// Experimental.
	SetPrimaryKeyField(val *string)
	// Experimental.
	PrimaryKeyFieldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextField() *string
	// Experimental.
	SetTextField(val *string)
	// Experimental.
	TextFieldInput() *string
	// Experimental.
	VectorField() *string
	// Experimental.
	SetVectorField(val *string)
	// Experimental.
	VectorFieldInput() *string
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
	ResetCustomMetadataField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference
type jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) CustomMetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMetadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) CustomMetadataFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMetadataFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) MetadataFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) PrimaryKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) PrimaryKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) TextFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) VectorField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) VectorFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorFieldInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsKnowledgeBase.StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference_Override(a AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsKnowledgeBase.StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetCustomMetadataField(val *string) {
	if err := j.validateSetCustomMetadataFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMetadataField",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetMetadataField(val *string) {
	if err := j.validateSetMetadataFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataField",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetPrimaryKeyField(val *string) {
	if err := j.validateSetPrimaryKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyField",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetTextField(val *string) {
	if err := j.validateSetTextFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textField",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference)SetVectorField(val *string) {
	if err := j.validateSetVectorFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorField",
		val,
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) ResetCustomMetadataField() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomMetadataField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

