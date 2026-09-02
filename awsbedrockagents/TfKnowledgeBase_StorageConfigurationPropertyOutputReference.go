package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKnowledgeBase_StorageConfigurationPropertyOutputReference interface {
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
	MongoDbAtlasConfiguration() TfKnowledgeBase_MongoDbAtlasConfigurationPropertyList
	// Experimental.
	MongoDbAtlasConfigurationInput() interface{}
	// Experimental.
	NeptuneAnalyticsConfiguration() TfKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList
	// Experimental.
	NeptuneAnalyticsConfigurationInput() interface{}
	// Experimental.
	OpensearchManagedClusterConfiguration() TfKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList
	// Experimental.
	OpensearchManagedClusterConfigurationInput() interface{}
	// Experimental.
	OpensearchServerlessConfiguration() TfKnowledgeBase_OpensearchServerlessConfigurationPropertyList
	// Experimental.
	OpensearchServerlessConfigurationInput() interface{}
	// Experimental.
	PineconeConfiguration() TfKnowledgeBase_PineconeConfigurationPropertyList
	// Experimental.
	PineconeConfigurationInput() interface{}
	// Experimental.
	RdsConfiguration() TfKnowledgeBase_RdsConfigurationPropertyList
	// Experimental.
	RdsConfigurationInput() interface{}
	// Experimental.
	RedisEnterpriseCloudConfiguration() TfKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList
	// Experimental.
	RedisEnterpriseCloudConfigurationInput() interface{}
	// Experimental.
	S3VectorsConfiguration() TfKnowledgeBase_S3VectorsConfigurationPropertyList
	// Experimental.
	S3VectorsConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutMongoDbAtlasConfiguration(value interface{})
	// Experimental.
	PutNeptuneAnalyticsConfiguration(value interface{})
	// Experimental.
	PutOpensearchManagedClusterConfiguration(value interface{})
	// Experimental.
	PutOpensearchServerlessConfiguration(value interface{})
	// Experimental.
	PutPineconeConfiguration(value interface{})
	// Experimental.
	PutRdsConfiguration(value interface{})
	// Experimental.
	PutRedisEnterpriseCloudConfiguration(value interface{})
	// Experimental.
	PutS3VectorsConfiguration(value interface{})
	// Experimental.
	ResetMongoDbAtlasConfiguration()
	// Experimental.
	ResetNeptuneAnalyticsConfiguration()
	// Experimental.
	ResetOpensearchManagedClusterConfiguration()
	// Experimental.
	ResetOpensearchServerlessConfiguration()
	// Experimental.
	ResetPineconeConfiguration()
	// Experimental.
	ResetRdsConfiguration()
	// Experimental.
	ResetRedisEnterpriseCloudConfiguration()
	// Experimental.
	ResetS3VectorsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfKnowledgeBase_StorageConfigurationPropertyOutputReference
type jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) MongoDbAtlasConfiguration() TfKnowledgeBase_MongoDbAtlasConfigurationPropertyList {
	var returns TfKnowledgeBase_MongoDbAtlasConfigurationPropertyList
	_jsii_.Get(
		j,
		"mongoDbAtlasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) MongoDbAtlasConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbAtlasConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) NeptuneAnalyticsConfiguration() TfKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList {
	var returns TfKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) NeptuneAnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchManagedClusterConfiguration() TfKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList {
	var returns TfKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchManagedClusterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchServerlessConfiguration() TfKnowledgeBase_OpensearchServerlessConfigurationPropertyList {
	var returns TfKnowledgeBase_OpensearchServerlessConfigurationPropertyList
	_jsii_.Get(
		j,
		"opensearchServerlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchServerlessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchServerlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PineconeConfiguration() TfKnowledgeBase_PineconeConfigurationPropertyList {
	var returns TfKnowledgeBase_PineconeConfigurationPropertyList
	_jsii_.Get(
		j,
		"pineconeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PineconeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pineconeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) RdsConfiguration() TfKnowledgeBase_RdsConfigurationPropertyList {
	var returns TfKnowledgeBase_RdsConfigurationPropertyList
	_jsii_.Get(
		j,
		"rdsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) RdsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) RedisEnterpriseCloudConfiguration() TfKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList {
	var returns TfKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) RedisEnterpriseCloudConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) S3VectorsConfiguration() TfKnowledgeBase_S3VectorsConfigurationPropertyList {
	var returns TfKnowledgeBase_S3VectorsConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3VectorsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) S3VectorsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3VectorsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfKnowledgeBase_StorageConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfKnowledgeBase_StorageConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfKnowledgeBase_StorageConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.StorageConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfKnowledgeBase_StorageConfigurationPropertyOutputReference_Override(t TfKnowledgeBase_StorageConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.StorageConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutMongoDbAtlasConfiguration(value interface{}) {
	if err := t.validatePutMongoDbAtlasConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMongoDbAtlasConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutNeptuneAnalyticsConfiguration(value interface{}) {
	if err := t.validatePutNeptuneAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNeptuneAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutOpensearchManagedClusterConfiguration(value interface{}) {
	if err := t.validatePutOpensearchManagedClusterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpensearchManagedClusterConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutOpensearchServerlessConfiguration(value interface{}) {
	if err := t.validatePutOpensearchServerlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpensearchServerlessConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutPineconeConfiguration(value interface{}) {
	if err := t.validatePutPineconeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPineconeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutRdsConfiguration(value interface{}) {
	if err := t.validatePutRdsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRdsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutRedisEnterpriseCloudConfiguration(value interface{}) {
	if err := t.validatePutRedisEnterpriseCloudConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedisEnterpriseCloudConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) PutS3VectorsConfiguration(value interface{}) {
	if err := t.validatePutS3VectorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3VectorsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetMongoDbAtlasConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMongoDbAtlasConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetNeptuneAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNeptuneAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetOpensearchManagedClusterConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetOpensearchManagedClusterConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetOpensearchServerlessConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetOpensearchServerlessConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetPineconeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetPineconeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetRdsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetRedisEnterpriseCloudConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetRedisEnterpriseCloudConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetS3VectorsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3VectorsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_StorageConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

