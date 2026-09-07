package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_KinesisSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_KinesisSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_KinesisSettingsProperty)
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

// The jsii proxy struct for AwsEndpoint_KinesisSettingsPropertyOutputReference
type jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) InternalValue() *AwsEndpoint_KinesisSettingsProperty {
	var returns *AwsEndpoint_KinesisSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) StreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) UseLargeIntegerValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLargeIntegerValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) UseLargeIntegerValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLargeIntegerValueInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_KinesisSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_KinesisSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_KinesisSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.KinesisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_KinesisSettingsPropertyOutputReference_Override(a AwsEndpoint_KinesisSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.KinesisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeControlDetails(val interface{}) {
	if err := j.validateSetIncludeControlDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeNullAndEmpty(val interface{}) {
	if err := j.validateSetIncludeNullAndEmptyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetIncludePartitionValue(val interface{}) {
	if err := j.validateSetIncludePartitionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeTableAlterOperations(val interface{}) {
	if err := j.validateSetIncludeTableAlterOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetIncludeTransactionDetails(val interface{}) {
	if err := j.validateSetIncludeTransactionDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_KinesisSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetMessageFormat(val *string) {
	if err := j.validateSetMessageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetPartitionIncludeSchemaTable(val interface{}) {
	if err := j.validateSetPartitionIncludeSchemaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetStreamArn(val *string) {
	if err := j.validateSetStreamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference)SetUseLargeIntegerValue(val interface{}) {
	if err := j.validateSetUseLargeIntegerValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLargeIntegerValue",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		a,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetStreamArn() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ResetUseLargeIntegerValue() {
	_jsii_.InvokeVoid(
		a,
		"resetUseLargeIntegerValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_KinesisSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

