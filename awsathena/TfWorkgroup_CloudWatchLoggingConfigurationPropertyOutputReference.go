package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsathena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsathena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWorkgroup_CloudWatchLoggingConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfWorkgroup_CloudWatchLoggingConfigurationProperty)
	// Experimental.
	LogGroup() *string
	// Experimental.
	SetLogGroup(val *string)
	// Experimental.
	LogGroupInput() *string
	// Experimental.
	LogStreamNamePrefix() *string
	// Experimental.
	SetLogStreamNamePrefix(val *string)
	// Experimental.
	LogStreamNamePrefixInput() *string
	// Experimental.
	LogType() TfWorkgroup_LogTypePropertyList
	// Experimental.
	LogTypeInput() interface{}
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
	PutLogType(value interface{})
	// Experimental.
	ResetLogGroup()
	// Experimental.
	ResetLogStreamNamePrefix()
	// Experimental.
	ResetLogType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference
type jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) InternalValue() *TfWorkgroup_CloudWatchLoggingConfigurationProperty {
	var returns *TfWorkgroup_CloudWatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogStreamNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogStreamNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogType() TfWorkgroup_LogTypePropertyList {
	var returns TfWorkgroup_LogTypePropertyList
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) LogTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.CloudWatchLoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference_Override(t TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.CloudWatchLoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetInternalValue(val *TfWorkgroup_CloudWatchLoggingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetLogGroup(val *string) {
	if err := j.validateSetLogGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetLogStreamNamePrefix(val *string) {
	if err := j.validateSetLogStreamNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamNamePrefix",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) PutLogType(value interface{}) {
	if err := t.validatePutLogTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ResetLogGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetLogGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ResetLogStreamNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetLogStreamNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ResetLogType() {
	_jsii_.InvokeVoid(
		t,
		"resetLogType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

