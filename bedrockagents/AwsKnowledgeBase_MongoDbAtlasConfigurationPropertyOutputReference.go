package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CollectionName() *string
	// Experimental.
	SetCollectionName(val *string)
	// Experimental.
	CollectionNameInput() *string
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
	CredentialsSecretArn() *string
	// Experimental.
	SetCredentialsSecretArn(val *string)
	// Experimental.
	CredentialsSecretArnInput() *string
	// Experimental.
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	Endpoint() *string
	// Experimental.
	SetEndpoint(val *string)
	// Experimental.
	EndpointInput() *string
	// Experimental.
	EndpointServiceName() *string
	// Experimental.
	SetEndpointServiceName(val *string)
	// Experimental.
	EndpointServiceNameInput() *string
	// Experimental.
	FieldMapping() AwsKnowledgeBase_StorageConfigurationMongoDbAtlasConfigurationFieldMappingPropertyList
	// Experimental.
	FieldMappingInput() interface{}
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
	TextIndexName() *string
	// Experimental.
	SetTextIndexName(val *string)
	// Experimental.
	TextIndexNameInput() *string
	// Experimental.
	VectorIndexName() *string
	// Experimental.
	SetVectorIndexName(val *string)
	// Experimental.
	VectorIndexNameInput() *string
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
	PutFieldMapping(value interface{})
	// Experimental.
	ResetEndpointServiceName()
	// Experimental.
	ResetFieldMapping()
	// Experimental.
	ResetTextIndexName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference
type jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) CollectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) CredentialsSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) EndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) EndpointServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) FieldMapping() AwsKnowledgeBase_StorageConfigurationMongoDbAtlasConfigurationFieldMappingPropertyList {
	var returns AwsKnowledgeBase_StorageConfigurationMongoDbAtlasConfigurationFieldMappingPropertyList
	_jsii_.Get(
		j,
		"fieldMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) FieldMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) TextIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textIndexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) TextIndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textIndexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) VectorIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) VectorIndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsKnowledgeBase.MongoDbAtlasConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference_Override(a AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsKnowledgeBase.MongoDbAtlasConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetCollectionName(val *string) {
	if err := j.validateSetCollectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionName",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetCredentialsSecretArn(val *string) {
	if err := j.validateSetCredentialsSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsSecretArn",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetEndpointServiceName(val *string) {
	if err := j.validateSetEndpointServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointServiceName",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetTextIndexName(val *string) {
	if err := j.validateSetTextIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textIndexName",
		val,
	)
}

func (j *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference)SetVectorIndexName(val *string) {
	if err := j.validateSetVectorIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorIndexName",
		val,
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) PutFieldMapping(value interface{}) {
	if err := a.validatePutFieldMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFieldMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ResetEndpointServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ResetFieldMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ResetTextIndexName() {
	_jsii_.InvokeVoid(
		a,
		"resetTextIndexName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKnowledgeBase_MongoDbAtlasConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

