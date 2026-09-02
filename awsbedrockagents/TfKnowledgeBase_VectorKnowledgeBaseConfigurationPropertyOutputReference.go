package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference interface {
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
	EmbeddingModelArn() *string
	// Experimental.
	SetEmbeddingModelArn(val *string)
	// Experimental.
	EmbeddingModelArnInput() *string
	// Experimental.
	EmbeddingModelConfiguration() TfKnowledgeBase_KnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationPropertyList
	// Experimental.
	EmbeddingModelConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SupplementalDataStorageConfiguration() TfKnowledgeBase_SupplementalDataStorageConfigurationPropertyList
	// Experimental.
	SupplementalDataStorageConfigurationInput() interface{}
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
	PutEmbeddingModelConfiguration(value interface{})
	// Experimental.
	PutSupplementalDataStorageConfiguration(value interface{})
	// Experimental.
	ResetEmbeddingModelConfiguration()
	// Experimental.
	ResetSupplementalDataStorageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference
type jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) EmbeddingModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) EmbeddingModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) EmbeddingModelConfiguration() TfKnowledgeBase_KnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationPropertyList {
	var returns TfKnowledgeBase_KnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationPropertyList
	_jsii_.Get(
		j,
		"embeddingModelConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) EmbeddingModelConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingModelConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) SupplementalDataStorageConfiguration() TfKnowledgeBase_SupplementalDataStorageConfigurationPropertyList {
	var returns TfKnowledgeBase_SupplementalDataStorageConfigurationPropertyList
	_jsii_.Get(
		j,
		"supplementalDataStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) SupplementalDataStorageConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supplementalDataStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.VectorKnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference_Override(t TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.VectorKnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetEmbeddingModelArn(val *string) {
	if err := j.validateSetEmbeddingModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModelArn",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) PutEmbeddingModelConfiguration(value interface{}) {
	if err := t.validatePutEmbeddingModelConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmbeddingModelConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) PutSupplementalDataStorageConfiguration(value interface{}) {
	if err := t.validatePutSupplementalDataStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSupplementalDataStorageConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ResetEmbeddingModelConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEmbeddingModelConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ResetSupplementalDataStorageConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSupplementalDataStorageConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

