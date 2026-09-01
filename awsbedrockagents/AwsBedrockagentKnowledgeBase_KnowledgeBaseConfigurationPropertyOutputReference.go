package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference interface {
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
	KendraKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList
	// Experimental.
	KendraKnowledgeBaseConfigurationInput() interface{}
	// Experimental.
	ManagedKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList
	// Experimental.
	ManagedKnowledgeBaseConfigurationInput() interface{}
	// Experimental.
	SqlKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList
	// Experimental.
	SqlKnowledgeBaseConfigurationInput() interface{}
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
	VectorKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList
	// Experimental.
	VectorKnowledgeBaseConfigurationInput() interface{}
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
	PutKendraKnowledgeBaseConfiguration(value interface{})
	// Experimental.
	PutManagedKnowledgeBaseConfiguration(value interface{})
	// Experimental.
	PutSqlKnowledgeBaseConfiguration(value interface{})
	// Experimental.
	PutVectorKnowledgeBaseConfiguration(value interface{})
	// Experimental.
	ResetKendraKnowledgeBaseConfiguration()
	// Experimental.
	ResetManagedKnowledgeBaseConfiguration()
	// Experimental.
	ResetSqlKnowledgeBaseConfiguration()
	// Experimental.
	ResetVectorKnowledgeBaseConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference
type jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) KendraKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"kendraKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) KendraKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kendraKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ManagedKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ManagedKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) SqlKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"sqlKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) SqlKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) VectorKnowledgeBaseConfiguration() AwsBedrockagentKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList {
	var returns AwsBedrockagentKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"vectorKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) VectorKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.KnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference_Override(a AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentKnowledgeBase.KnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutKendraKnowledgeBaseConfiguration(value interface{}) {
	if err := a.validatePutKendraKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKendraKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutManagedKnowledgeBaseConfiguration(value interface{}) {
	if err := a.validatePutManagedKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutSqlKnowledgeBaseConfiguration(value interface{}) {
	if err := a.validatePutSqlKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutVectorKnowledgeBaseConfiguration(value interface{}) {
	if err := a.validatePutVectorKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVectorKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetKendraKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKendraKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetManagedKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetSqlKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetVectorKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVectorKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

