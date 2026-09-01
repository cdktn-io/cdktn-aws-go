package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference interface {
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
	ResetMetadataField()
	// Experimental.
	ResetTextField()
	// Experimental.
	ResetVectorField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference
type jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) MetadataFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) TextFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) VectorField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) VectorFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorFieldInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference_Override(a AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetMetadataField(val *string) {
	if err := j.validateSetMetadataFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataField",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetTextField(val *string) {
	if err := j.validateSetTextFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textField",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference)SetVectorField(val *string) {
	if err := j.validateSetVectorFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorField",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ResetMetadataField() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ResetTextField() {
	_jsii_.InvokeVoid(
		a,
		"resetTextField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ResetVectorField() {
	_jsii_.InvokeVoid(
		a,
		"resetVectorField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

