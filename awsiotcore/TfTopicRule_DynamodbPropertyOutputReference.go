package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicRule_DynamodbPropertyOutputReference interface {
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
	HashKeyField() *string
	// Experimental.
	SetHashKeyField(val *string)
	// Experimental.
	HashKeyFieldInput() *string
	// Experimental.
	HashKeyType() *string
	// Experimental.
	SetHashKeyType(val *string)
	// Experimental.
	HashKeyTypeInput() *string
	// Experimental.
	HashKeyValue() *string
	// Experimental.
	SetHashKeyValue(val *string)
	// Experimental.
	HashKeyValueInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Operation() *string
	// Experimental.
	SetOperation(val *string)
	// Experimental.
	OperationInput() *string
	// Experimental.
	PayloadField() *string
	// Experimental.
	SetPayloadField(val *string)
	// Experimental.
	PayloadFieldInput() *string
	// Experimental.
	RangeKeyField() *string
	// Experimental.
	SetRangeKeyField(val *string)
	// Experimental.
	RangeKeyFieldInput() *string
	// Experimental.
	RangeKeyType() *string
	// Experimental.
	SetRangeKeyType(val *string)
	// Experimental.
	RangeKeyTypeInput() *string
	// Experimental.
	RangeKeyValue() *string
	// Experimental.
	SetRangeKeyValue(val *string)
	// Experimental.
	RangeKeyValueInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	TableName() *string
	// Experimental.
	SetTableName(val *string)
	// Experimental.
	TableNameInput() *string
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
	ResetHashKeyType()
	// Experimental.
	ResetOperation()
	// Experimental.
	ResetPayloadField()
	// Experimental.
	ResetRangeKeyField()
	// Experimental.
	ResetRangeKeyType()
	// Experimental.
	ResetRangeKeyValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTopicRule_DynamodbPropertyOutputReference
type jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) HashKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) PayloadField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) PayloadFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RangeKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTopicRule_DynamodbPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTopicRule_DynamodbPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTopicRule_DynamodbPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.DynamodbPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTopicRule_DynamodbPropertyOutputReference_Override(t TfTopicRule_DynamodbPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.DynamodbPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetHashKeyField(val *string) {
	if err := j.validateSetHashKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyField",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetHashKeyType(val *string) {
	if err := j.validateSetHashKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyType",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetHashKeyValue(val *string) {
	if err := j.validateSetHashKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetOperation(val *string) {
	if err := j.validateSetOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetPayloadField(val *string) {
	if err := j.validateSetPayloadFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadField",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetRangeKeyField(val *string) {
	if err := j.validateSetRangeKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyField",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetRangeKeyType(val *string) {
	if err := j.validateSetRangeKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyType",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetRangeKeyValue(val *string) {
	if err := j.validateSetRangeKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetHashKeyType() {
	_jsii_.InvokeVoid(
		t,
		"resetHashKeyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetPayloadField() {
	_jsii_.InvokeVoid(
		t,
		"resetPayloadField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetRangeKeyField() {
	_jsii_.InvokeVoid(
		t,
		"resetRangeKeyField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetRangeKeyType() {
	_jsii_.InvokeVoid(
		t,
		"resetRangeKeyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ResetRangeKeyValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRangeKeyValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTopicRule_DynamodbPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

