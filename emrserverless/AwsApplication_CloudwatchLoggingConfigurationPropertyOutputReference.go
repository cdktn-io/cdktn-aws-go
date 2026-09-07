package emrserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrserverless/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference interface {
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
	EncryptionKeyArn() *string
	// Experimental.
	SetEncryptionKeyArn(val *string)
	// Experimental.
	EncryptionKeyArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsApplication_CloudwatchLoggingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_CloudwatchLoggingConfigurationProperty)
	// Experimental.
	LogGroupName() *string
	// Experimental.
	SetLogGroupName(val *string)
	// Experimental.
	LogGroupNameInput() *string
	// Experimental.
	LogStreamNamePrefix() *string
	// Experimental.
	SetLogStreamNamePrefix(val *string)
	// Experimental.
	LogStreamNamePrefixInput() *string
	// Experimental.
	LogTypes() AwsApplication_LogTypesPropertyList
	// Experimental.
	LogTypesInput() interface{}
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
	PutLogTypes(value interface{})
	// Experimental.
	ResetEncryptionKeyArn()
	// Experimental.
	ResetLogGroupName()
	// Experimental.
	ResetLogStreamNamePrefix()
	// Experimental.
	ResetLogTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference
type jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) EncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) InternalValue() *AwsApplication_CloudwatchLoggingConfigurationProperty {
	var returns *AwsApplication_CloudwatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogStreamNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogStreamNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogTypes() AwsApplication_LogTypesPropertyList {
	var returns AwsApplication_LogTypesPropertyList
	_jsii_.Get(
		j,
		"logTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) LogTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_CloudwatchLoggingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsApplication.CloudwatchLoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference_Override(a AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsApplication.CloudwatchLoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetEncryptionKeyArn(val *string) {
	if err := j.validateSetEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetInternalValue(val *AwsApplication_CloudwatchLoggingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetLogStreamNamePrefix(val *string) {
	if err := j.validateSetLogStreamNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamNamePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) PutLogTypes(value interface{}) {
	if err := a.validatePutLogTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogTypes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ResetEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ResetLogStreamNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetLogStreamNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ResetLogTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetLogTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

