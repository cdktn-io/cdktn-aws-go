package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference interface {
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
	MongoDbAtlasConfiguration() AwsBedrockagentKnowledgeBase_MongoDbAtlasConfigurationPropertyList
	// Experimental.
	MongoDbAtlasConfigurationInput() interface{}
	// Experimental.
	NeptuneAnalyticsConfiguration() AwsBedrockagentKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList
	// Experimental.
	NeptuneAnalyticsConfigurationInput() interface{}
	// Experimental.
	OpensearchManagedClusterConfiguration() AwsBedrockagentKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList
	// Experimental.
	OpensearchManagedClusterConfigurationInput() interface{}
	// Experimental.
	OpensearchServerlessConfiguration() AwsBedrockagentKnowledgeBase_OpensearchServerlessConfigurationPropertyList
	// Experimental.
	OpensearchServerlessConfigurationInput() interface{}
	// Experimental.
	PineconeConfiguration() AwsBedrockagentKnowledgeBase_PineconeConfigurationPropertyList
	// Experimental.
	PineconeConfigurationInput() interface{}
	// Experimental.
	RdsConfiguration() AwsBedrockagentKnowledgeBase_RdsConfigurationPropertyList
	// Experimental.
	RdsConfigurationInput() interface{}
	// Experimental.
	RedisEnterpriseCloudConfiguration() AwsBedrockagentKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList
	// Experimental.
	RedisEnterpriseCloudConfigurationInput() interface{}
	// Experimental.
	S3VectorsConfiguration() AwsBedrockagentKnowledgeBase_S3VectorsConfigurationPropertyList
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

// The jsii proxy struct for AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference
type jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) MongoDbAtlasConfiguration() AwsBedrockagentKnowledgeBase_MongoDbAtlasConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_MongoDbAtlasConfigurationPropertyList
	_jsii_.Get(
		j,
		"mongoDbAtlasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) MongoDbAtlasConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbAtlasConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) NeptuneAnalyticsConfiguration() AwsBedrockagentKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_NeptuneAnalyticsConfigurationPropertyList
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) NeptuneAnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchManagedClusterConfiguration() AwsBedrockagentKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_OpensearchManagedClusterConfigurationPropertyList
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchManagedClusterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchServerlessConfiguration() AwsBedrockagentKnowledgeBase_OpensearchServerlessConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_OpensearchServerlessConfigurationPropertyList
	_jsii_.Get(
		j,
		"opensearchServerlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) OpensearchServerlessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchServerlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PineconeConfiguration() AwsBedrockagentKnowledgeBase_PineconeConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_PineconeConfigurationPropertyList
	_jsii_.Get(
		j,
		"pineconeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PineconeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pineconeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) RdsConfiguration() AwsBedrockagentKnowledgeBase_RdsConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_RdsConfigurationPropertyList
	_jsii_.Get(
		j,
		"rdsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) RdsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) RedisEnterpriseCloudConfiguration() AwsBedrockagentKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_RedisEnterpriseCloudConfigurationPropertyList
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) RedisEnterpriseCloudConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) S3VectorsConfiguration() AwsBedrockagentKnowledgeBase_S3VectorsConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_S3VectorsConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3VectorsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) S3VectorsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3VectorsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.StorageConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference_Override(a AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.StorageConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutMongoDbAtlasConfiguration(value interface{}) {
	if err := a.validatePutMongoDbAtlasConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMongoDbAtlasConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutNeptuneAnalyticsConfiguration(value interface{}) {
	if err := a.validatePutNeptuneAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNeptuneAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutOpensearchManagedClusterConfiguration(value interface{}) {
	if err := a.validatePutOpensearchManagedClusterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchManagedClusterConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutOpensearchServerlessConfiguration(value interface{}) {
	if err := a.validatePutOpensearchServerlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchServerlessConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutPineconeConfiguration(value interface{}) {
	if err := a.validatePutPineconeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPineconeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutRdsConfiguration(value interface{}) {
	if err := a.validatePutRdsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutRedisEnterpriseCloudConfiguration(value interface{}) {
	if err := a.validatePutRedisEnterpriseCloudConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedisEnterpriseCloudConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) PutS3VectorsConfiguration(value interface{}) {
	if err := a.validatePutS3VectorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3VectorsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetMongoDbAtlasConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMongoDbAtlasConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetNeptuneAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNeptuneAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetOpensearchManagedClusterConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchManagedClusterConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetOpensearchServerlessConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchServerlessConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetPineconeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPineconeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetRdsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetRedisEnterpriseCloudConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRedisEnterpriseCloudConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ResetS3VectorsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3VectorsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_StorageConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

