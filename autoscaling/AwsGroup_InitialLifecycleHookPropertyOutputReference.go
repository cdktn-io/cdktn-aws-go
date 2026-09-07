package autoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/autoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/autoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGroup_InitialLifecycleHookPropertyOutputReference interface {
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
	DefaultResult() *string
	// Experimental.
	SetDefaultResult(val *string)
	// Experimental.
	DefaultResultInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HeartbeatTimeout() *float64
	// Experimental.
	SetHeartbeatTimeout(val *float64)
	// Experimental.
	HeartbeatTimeoutInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LifecycleTransition() *string
	// Experimental.
	SetLifecycleTransition(val *string)
	// Experimental.
	LifecycleTransitionInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NotificationMetadata() *string
	// Experimental.
	SetNotificationMetadata(val *string)
	// Experimental.
	NotificationMetadataInput() *string
	// Experimental.
	NotificationTargetArn() *string
	// Experimental.
	SetNotificationTargetArn(val *string)
	// Experimental.
	NotificationTargetArnInput() *string
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
	ResetDefaultResult()
	// Experimental.
	ResetHeartbeatTimeout()
	// Experimental.
	ResetNotificationMetadata()
	// Experimental.
	ResetNotificationTargetArn()
	// Experimental.
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGroup_InitialLifecycleHookPropertyOutputReference
type jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) DefaultResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) DefaultResultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) HeartbeatTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) HeartbeatTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) LifecycleTransition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) LifecycleTransitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) NotificationMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) NotificationMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) NotificationTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) NotificationTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGroup_InitialLifecycleHookPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGroup_InitialLifecycleHookPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGroup_InitialLifecycleHookPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup.InitialLifecycleHookPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGroup_InitialLifecycleHookPropertyOutputReference_Override(a AwsGroup_InitialLifecycleHookPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup.InitialLifecycleHookPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetDefaultResult(val *string) {
	if err := j.validateSetDefaultResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultResult",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetHeartbeatTimeout(val *float64) {
	if err := j.validateSetHeartbeatTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetLifecycleTransition(val *string) {
	if err := j.validateSetLifecycleTransitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleTransition",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetNotificationMetadata(val *string) {
	if err := j.validateSetNotificationMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationMetadata",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetNotificationTargetArn(val *string) {
	if err := j.validateSetNotificationTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTargetArn",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ResetDefaultResult() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResult",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ResetHeartbeatTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ResetNotificationMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ResetNotificationTargetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTargetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGroup_InitialLifecycleHookPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

