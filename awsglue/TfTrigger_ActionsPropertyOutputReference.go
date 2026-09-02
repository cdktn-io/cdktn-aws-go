package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrigger_ActionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Arguments() *map[string]*string
	// Experimental.
	SetArguments(val *map[string]*string)
	// Experimental.
	ArgumentsInput() *map[string]*string
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
	// Experimental.
	CrawlerName() *string
	// Experimental.
	SetCrawlerName(val *string)
	// Experimental.
	CrawlerNameInput() *string
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
	JobName() *string
	// Experimental.
	SetJobName(val *string)
	// Experimental.
	JobNameInput() *string
	// Experimental.
	NotificationProperty() TfTrigger_NotificationPropertyPropertyOutputReference
	// Experimental.
	NotificationPropertyInput() *TfTrigger_NotificationPropertyProperty
	// Experimental.
	SecurityConfiguration() *string
	// Experimental.
	SetSecurityConfiguration(val *string)
	// Experimental.
	SecurityConfigurationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() *float64
	// Experimental.
	SetTimeout(val *float64)
	// Experimental.
	TimeoutInput() *float64
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
	PutNotificationProperty(value *TfTrigger_NotificationPropertyProperty)
	// Experimental.
	ResetArguments()
	// Experimental.
	ResetCrawlerName()
	// Experimental.
	ResetJobName()
	// Experimental.
	ResetNotificationProperty()
	// Experimental.
	ResetSecurityConfiguration()
	// Experimental.
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTrigger_ActionsPropertyOutputReference
type jsiiProxy_TfTrigger_ActionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) Arguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) CrawlerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) CrawlerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) NotificationProperty() TfTrigger_NotificationPropertyPropertyOutputReference {
	var returns TfTrigger_NotificationPropertyPropertyOutputReference
	_jsii_.Get(
		j,
		"notificationProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) NotificationPropertyInput() *TfTrigger_NotificationPropertyProperty {
	var returns *TfTrigger_NotificationPropertyProperty
	_jsii_.Get(
		j,
		"notificationPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTrigger_ActionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTrigger_ActionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTrigger_ActionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrigger_ActionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfTrigger.ActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTrigger_ActionsPropertyOutputReference_Override(t TfTrigger_ActionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfTrigger.ActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetArguments(val *map[string]*string) {
	if err := j.validateSetArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetCrawlerName(val *string) {
	if err := j.validateSetCrawlerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerName",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTrigger_ActionsPropertyOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) PutNotificationProperty(value *TfTrigger_NotificationPropertyProperty) {
	if err := t.validatePutNotificationPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotificationProperty",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetCrawlerName() {
	_jsii_.InvokeVoid(
		t,
		"resetCrawlerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetJobName() {
	_jsii_.InvokeVoid(
		t,
		"resetJobName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetNotificationProperty() {
	_jsii_.InvokeVoid(
		t,
		"resetNotificationProperty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTrigger_ActionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

