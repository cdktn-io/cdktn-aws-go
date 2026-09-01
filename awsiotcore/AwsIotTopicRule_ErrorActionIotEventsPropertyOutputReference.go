package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchMode() interface{}
	// Experimental.
	SetBatchMode(val interface{})
	// Experimental.
	BatchModeInput() interface{}
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
	InputName() *string
	// Experimental.
	SetInputName(val *string)
	// Experimental.
	InputNameInput() *string
	// Experimental.
	InternalValue() *AwsIotTopicRule_ErrorActionIotEventsProperty
	// Experimental.
	SetInternalValue(val *AwsIotTopicRule_ErrorActionIotEventsProperty)
	// Experimental.
	MessageId() *string
	// Experimental.
	SetMessageId(val *string)
	// Experimental.
	MessageIdInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	ResetBatchMode()
	// Experimental.
	ResetMessageId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference
type jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) BatchMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) BatchModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) InputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) InternalValue() *AwsIotTopicRule_ErrorActionIotEventsProperty {
	var returns *AwsIotTopicRule_ErrorActionIotEventsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) MessageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) MessageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIotTopicRule_ErrorActionIotEventsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule.ErrorActionIotEventsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference_Override(a AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule.ErrorActionIotEventsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetBatchMode(val interface{}) {
	if err := j.validateSetBatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchMode",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetInputName(val *string) {
	if err := j.validateSetInputNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputName",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetInternalValue(val *AwsIotTopicRule_ErrorActionIotEventsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetMessageId(val *string) {
	if err := j.validateSetMessageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageId",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ResetBatchMode() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ResetMessageId() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

