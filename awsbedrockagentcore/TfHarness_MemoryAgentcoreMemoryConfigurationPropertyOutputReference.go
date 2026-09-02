package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActorId() *string
	// Experimental.
	SetActorId(val *string)
	// Experimental.
	ActorIdInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
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
	MessagesCount() *float64
	// Experimental.
	SetMessagesCount(val *float64)
	// Experimental.
	MessagesCountInput() *float64
	// Experimental.
	RetrievalConfig() TfHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyList
	// Experimental.
	RetrievalConfigInput() interface{}
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
	PutRetrievalConfig(value interface{})
	// Experimental.
	ResetActorId()
	// Experimental.
	ResetMessagesCount()
	// Experimental.
	ResetRetrievalConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference
type jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ActorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ActorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) MessagesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) MessagesCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) RetrievalConfig() TfHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyList {
	var returns TfHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyList
	_jsii_.Get(
		j,
		"retrievalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) RetrievalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrievalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfHarness.MemoryAgentcoreMemoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference_Override(t TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfHarness.MemoryAgentcoreMemoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetActorId(val *string) {
	if err := j.validateSetActorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actorId",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetMessagesCount(val *float64) {
	if err := j.validateSetMessagesCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagesCount",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) PutRetrievalConfig(value interface{}) {
	if err := t.validatePutRetrievalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetrievalConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ResetActorId() {
	_jsii_.InvokeVoid(
		t,
		"resetActorId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ResetMessagesCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMessagesCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ResetRetrievalConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRetrievalConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHarness_MemoryAgentcoreMemoryConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

