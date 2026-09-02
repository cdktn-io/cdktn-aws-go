package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference interface {
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
	KendraKnowledgeBaseConfiguration() TfKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList
	// Experimental.
	KendraKnowledgeBaseConfigurationInput() interface{}
	// Experimental.
	ManagedKnowledgeBaseConfiguration() TfKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList
	// Experimental.
	ManagedKnowledgeBaseConfigurationInput() interface{}
	// Experimental.
	SqlKnowledgeBaseConfiguration() TfKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList
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
	VectorKnowledgeBaseConfiguration() TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList
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

// The jsii proxy struct for TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference
type jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) KendraKnowledgeBaseConfiguration() TfKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList {
	var returns TfKnowledgeBase_KendraKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"kendraKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) KendraKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kendraKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ManagedKnowledgeBaseConfiguration() TfKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList {
	var returns TfKnowledgeBase_ManagedKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ManagedKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) SqlKnowledgeBaseConfiguration() TfKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList {
	var returns TfKnowledgeBase_SqlKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"sqlKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) SqlKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) VectorKnowledgeBaseConfiguration() TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList {
	var returns TfKnowledgeBase_VectorKnowledgeBaseConfigurationPropertyList
	_jsii_.Get(
		j,
		"vectorKnowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) VectorKnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorKnowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.KnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference_Override(t TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfKnowledgeBase.KnowledgeBaseConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutKendraKnowledgeBaseConfiguration(value interface{}) {
	if err := t.validatePutKendraKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKendraKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutManagedKnowledgeBaseConfiguration(value interface{}) {
	if err := t.validatePutManagedKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutSqlKnowledgeBaseConfiguration(value interface{}) {
	if err := t.validatePutSqlKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqlKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) PutVectorKnowledgeBaseConfiguration(value interface{}) {
	if err := t.validatePutVectorKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVectorKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetKendraKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetKendraKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetManagedKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetSqlKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSqlKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ResetVectorKnowledgeBaseConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVectorKnowledgeBaseConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfKnowledgeBase_KnowledgeBaseConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

