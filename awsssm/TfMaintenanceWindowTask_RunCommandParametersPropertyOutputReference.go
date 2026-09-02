package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchConfig() TfMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference
	// Experimental.
	CloudwatchConfigInput() *TfMaintenanceWindowTask_CloudwatchConfigProperty
	// Experimental.
	Comment() *string
	// Experimental.
	SetComment(val *string)
	// Experimental.
	CommentInput() *string
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
	DocumentHash() *string
	// Experimental.
	SetDocumentHash(val *string)
	// Experimental.
	DocumentHashInput() *string
	// Experimental.
	DocumentHashType() *string
	// Experimental.
	SetDocumentHashType(val *string)
	// Experimental.
	DocumentHashTypeInput() *string
	// Experimental.
	DocumentVersion() *string
	// Experimental.
	SetDocumentVersion(val *string)
	// Experimental.
	DocumentVersionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfMaintenanceWindowTask_RunCommandParametersProperty
	// Experimental.
	SetInternalValue(val *TfMaintenanceWindowTask_RunCommandParametersProperty)
	// Experimental.
	NotificationConfig() TfMaintenanceWindowTask_NotificationConfigPropertyOutputReference
	// Experimental.
	NotificationConfigInput() *TfMaintenanceWindowTask_NotificationConfigProperty
	// Experimental.
	OutputS3Bucket() *string
	// Experimental.
	SetOutputS3Bucket(val *string)
	// Experimental.
	OutputS3BucketInput() *string
	// Experimental.
	OutputS3KeyPrefix() *string
	// Experimental.
	SetOutputS3KeyPrefix(val *string)
	// Experimental.
	OutputS3KeyPrefixInput() *string
	// Experimental.
	Parameter() TfMaintenanceWindowTask_TaskInvocationParametersRunCommandParametersParameterPropertyList
	// Experimental.
	ParameterInput() interface{}
	// Experimental.
	ServiceRoleArn() *string
	// Experimental.
	SetServiceRoleArn(val *string)
	// Experimental.
	ServiceRoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutSeconds() *float64
	// Experimental.
	SetTimeoutSeconds(val *float64)
	// Experimental.
	TimeoutSecondsInput() *float64
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
	PutCloudwatchConfig(value *TfMaintenanceWindowTask_CloudwatchConfigProperty)
	// Experimental.
	PutNotificationConfig(value *TfMaintenanceWindowTask_NotificationConfigProperty)
	// Experimental.
	PutParameter(value interface{})
	// Experimental.
	ResetCloudwatchConfig()
	// Experimental.
	ResetComment()
	// Experimental.
	ResetDocumentHash()
	// Experimental.
	ResetDocumentHashType()
	// Experimental.
	ResetDocumentVersion()
	// Experimental.
	ResetNotificationConfig()
	// Experimental.
	ResetOutputS3Bucket()
	// Experimental.
	ResetOutputS3KeyPrefix()
	// Experimental.
	ResetParameter()
	// Experimental.
	ResetServiceRoleArn()
	// Experimental.
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference
type jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) CloudwatchConfig() TfMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference {
	var returns TfMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) CloudwatchConfigInput() *TfMaintenanceWindowTask_CloudwatchConfigProperty {
	var returns *TfMaintenanceWindowTask_CloudwatchConfigProperty
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentHashType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentHashTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) InternalValue() *TfMaintenanceWindowTask_RunCommandParametersProperty {
	var returns *TfMaintenanceWindowTask_RunCommandParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) NotificationConfig() TfMaintenanceWindowTask_NotificationConfigPropertyOutputReference {
	var returns TfMaintenanceWindowTask_NotificationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) NotificationConfigInput() *TfMaintenanceWindowTask_NotificationConfigProperty {
	var returns *TfMaintenanceWindowTask_NotificationConfigProperty
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) OutputS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) OutputS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) OutputS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) OutputS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) Parameter() TfMaintenanceWindowTask_TaskInvocationParametersRunCommandParametersParameterPropertyList {
	var returns TfMaintenanceWindowTask_TaskInvocationParametersRunCommandParametersParameterPropertyList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMaintenanceWindowTask_RunCommandParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask.RunCommandParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference_Override(t TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask.RunCommandParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetDocumentHash(val *string) {
	if err := j.validateSetDocumentHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentHash",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetDocumentHashType(val *string) {
	if err := j.validateSetDocumentHashTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentHashType",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetInternalValue(val *TfMaintenanceWindowTask_RunCommandParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetOutputS3Bucket(val *string) {
	if err := j.validateSetOutputS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3Bucket",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetOutputS3KeyPrefix(val *string) {
	if err := j.validateSetOutputS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) PutCloudwatchConfig(value *TfMaintenanceWindowTask_CloudwatchConfigProperty) {
	if err := t.validatePutCloudwatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) PutNotificationConfig(value *TfMaintenanceWindowTask_NotificationConfigProperty) {
	if err := t.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) PutParameter(value interface{}) {
	if err := t.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParameter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		t,
		"resetComment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetDocumentHash() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentHash",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetDocumentHashType() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentHashType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetOutputS3Bucket() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputS3Bucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetOutputS3KeyPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputS3KeyPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		t,
		"resetParameter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

