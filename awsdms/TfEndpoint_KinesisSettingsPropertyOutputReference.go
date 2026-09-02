package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_KinesisSettingsPropertyOutputReference interface {
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
	IncludeControlDetails() interface{}
	// Experimental.
	SetIncludeControlDetails(val interface{})
	// Experimental.
	IncludeControlDetailsInput() interface{}
	// Experimental.
	IncludeNullAndEmpty() interface{}
	// Experimental.
	SetIncludeNullAndEmpty(val interface{})
	// Experimental.
	IncludeNullAndEmptyInput() interface{}
	// Experimental.
	IncludePartitionValue() interface{}
	// Experimental.
	SetIncludePartitionValue(val interface{})
	// Experimental.
	IncludePartitionValueInput() interface{}
	// Experimental.
	IncludeTableAlterOperations() interface{}
	// Experimental.
	SetIncludeTableAlterOperations(val interface{})
	// Experimental.
	IncludeTableAlterOperationsInput() interface{}
	// Experimental.
	IncludeTransactionDetails() interface{}
	// Experimental.
	SetIncludeTransactionDetails(val interface{})
	// Experimental.
	IncludeTransactionDetailsInput() interface{}
	// Experimental.
	InternalValue() *TfEndpoint_KinesisSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_KinesisSettingsProperty)
	// Experimental.
	MessageFormat() *string
	// Experimental.
	SetMessageFormat(val *string)
	// Experimental.
	MessageFormatInput() *string
	// Experimental.
	PartitionIncludeSchemaTable() interface{}
	// Experimental.
	SetPartitionIncludeSchemaTable(val interface{})
	// Experimental.
	PartitionIncludeSchemaTableInput() interface{}
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	SetServiceAccessRoleArn(val *string)
	// Experimental.
	ServiceAccessRoleArnInput() *string
	// Experimental.
	StreamArn() *string
	// Experimental.
	SetStreamArn(val *string)
	// Experimental.
	StreamArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseLargeIntegerValue() interface{}
	// Experimental.
	SetUseLargeIntegerValue(val interface{})
	// Experimental.
	UseLargeIntegerValueInput() interface{}
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
	ResetIncludeControlDetails()
	// Experimental.
	ResetIncludeNullAndEmpty()
	// Experimental.
	ResetIncludePartitionValue()
	// Experimental.
	ResetIncludeTableAlterOperations()
	// Experimental.
	ResetIncludeTransactionDetails()
	// Experimental.
	ResetMessageFormat()
	// Experimental.
	ResetPartitionIncludeSchemaTable()
	// Experimental.
	ResetServiceAccessRoleArn()
	// Experimental.
	ResetStreamArn()
	// Experimental.
	ResetUseLargeIntegerValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_KinesisSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) InternalValue() *TfEndpoint_KinesisSettingsProperty {
	var returns *TfEndpoint_KinesisSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) StreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) UseLargeIntegerValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLargeIntegerValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) UseLargeIntegerValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLargeIntegerValueInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_KinesisSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_KinesisSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_KinesisSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.KinesisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_KinesisSettingsPropertyOutputReference_Override(t TfEndpoint_KinesisSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.KinesisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeControlDetails(val interface{}) {
	if err := j.validateSetIncludeControlDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeNullAndEmpty(val interface{}) {
	if err := j.validateSetIncludeNullAndEmptyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetIncludePartitionValue(val interface{}) {
	if err := j.validateSetIncludePartitionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeTableAlterOperations(val interface{}) {
	if err := j.validateSetIncludeTableAlterOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeTransactionDetails(val interface{}) {
	if err := j.validateSetIncludeTransactionDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_KinesisSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetMessageFormat(val *string) {
	if err := j.validateSetMessageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetPartitionIncludeSchemaTable(val interface{}) {
	if err := j.validateSetPartitionIncludeSchemaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetStreamArn(val *string) {
	if err := j.validateSetStreamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference)SetUseLargeIntegerValue(val interface{}) {
	if err := j.validateSetUseLargeIntegerValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLargeIntegerValue",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		t,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetStreamArn() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ResetUseLargeIntegerValue() {
	_jsii_.InvokeVoid(
		t,
		"resetUseLargeIntegerValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_KinesisSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

