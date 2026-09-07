package backup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/backup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/backup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPlan_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CompletionWindow() *float64
	// Experimental.
	SetCompletionWindow(val *float64)
	// Experimental.
	CompletionWindowInput() *float64
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
	CopyAction() AwsPlan_CopyActionPropertyList
	// Experimental.
	CopyActionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableContinuousBackup() interface{}
	// Experimental.
	SetEnableContinuousBackup(val interface{})
	// Experimental.
	EnableContinuousBackupInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Lifecycle() AwsPlan_RuleLifecyclePropertyOutputReference
	// Experimental.
	LifecycleInput() *AwsPlan_RuleLifecycleProperty
	// Experimental.
	RecoveryPointTags() *map[string]*string
	// Experimental.
	SetRecoveryPointTags(val *map[string]*string)
	// Experimental.
	RecoveryPointTagsInput() *map[string]*string
	// Experimental.
	RuleName() *string
	// Experimental.
	SetRuleName(val *string)
	// Experimental.
	RuleNameInput() *string
	// Experimental.
	ScanAction() AwsPlan_ScanActionPropertyList
	// Experimental.
	ScanActionInput() interface{}
	// Experimental.
	Schedule() *string
	// Experimental.
	SetSchedule(val *string)
	// Experimental.
	ScheduleExpressionTimezone() *string
	// Experimental.
	SetScheduleExpressionTimezone(val *string)
	// Experimental.
	ScheduleExpressionTimezoneInput() *string
	// Experimental.
	ScheduleInput() *string
	// Experimental.
	StartWindow() *float64
	// Experimental.
	SetStartWindow(val *float64)
	// Experimental.
	StartWindowInput() *float64
	// Experimental.
	TargetLogicallyAirGappedBackupVaultArn() *string
	// Experimental.
	SetTargetLogicallyAirGappedBackupVaultArn(val *string)
	// Experimental.
	TargetLogicallyAirGappedBackupVaultArnInput() *string
	// Experimental.
	TargetVaultName() *string
	// Experimental.
	SetTargetVaultName(val *string)
	// Experimental.
	TargetVaultNameInput() *string
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
	PutCopyAction(value interface{})
	// Experimental.
	PutLifecycle(value *AwsPlan_RuleLifecycleProperty)
	// Experimental.
	PutScanAction(value interface{})
	// Experimental.
	ResetCompletionWindow()
	// Experimental.
	ResetCopyAction()
	// Experimental.
	ResetEnableContinuousBackup()
	// Experimental.
	ResetLifecycle()
	// Experimental.
	ResetRecoveryPointTags()
	// Experimental.
	ResetScanAction()
	// Experimental.
	ResetSchedule()
	// Experimental.
	ResetScheduleExpressionTimezone()
	// Experimental.
	ResetStartWindow()
	// Experimental.
	ResetTargetLogicallyAirGappedBackupVaultArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPlan_RulePropertyOutputReference
type jsiiProxy_AwsPlan_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) CompletionWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) CompletionWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) CopyAction() AwsPlan_CopyActionPropertyList {
	var returns AwsPlan_CopyActionPropertyList
	_jsii_.Get(
		j,
		"copyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) CopyActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) EnableContinuousBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) EnableContinuousBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) Lifecycle() AwsPlan_RuleLifecyclePropertyOutputReference {
	var returns AwsPlan_RuleLifecyclePropertyOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) LifecycleInput() *AwsPlan_RuleLifecycleProperty {
	var returns *AwsPlan_RuleLifecycleProperty
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) RecoveryPointTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) RecoveryPointTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ScanAction() AwsPlan_ScanActionPropertyList {
	var returns AwsPlan_ScanActionPropertyList
	_jsii_.Get(
		j,
		"scanAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ScanActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scanActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ScheduleExpressionTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ScheduleExpressionTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) StartWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) StartWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TargetLogicallyAirGappedBackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetLogicallyAirGappedBackupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TargetLogicallyAirGappedBackupVaultArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetLogicallyAirGappedBackupVaultArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TargetVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TargetVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPlan_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPlan_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPlan_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPlan_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.AwsPlan.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPlan_RulePropertyOutputReference_Override(a AwsPlan_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.AwsPlan.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetCompletionWindow(val *float64) {
	if err := j.validateSetCompletionWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completionWindow",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetEnableContinuousBackup(val interface{}) {
	if err := j.validateSetEnableContinuousBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableContinuousBackup",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetRecoveryPointTags(val *map[string]*string) {
	if err := j.validateSetRecoveryPointTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTags",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetScheduleExpressionTimezone(val *string) {
	if err := j.validateSetScheduleExpressionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpressionTimezone",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetStartWindow(val *float64) {
	if err := j.validateSetStartWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startWindow",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetTargetLogicallyAirGappedBackupVaultArn(val *string) {
	if err := j.validateSetTargetLogicallyAirGappedBackupVaultArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetLogicallyAirGappedBackupVaultArn",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetTargetVaultName(val *string) {
	if err := j.validateSetTargetVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVaultName",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) PutCopyAction(value interface{}) {
	if err := a.validatePutCopyActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCopyAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) PutLifecycle(value *AwsPlan_RuleLifecycleProperty) {
	if err := a.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) PutScanAction(value interface{}) {
	if err := a.validatePutScanActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScanAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetCompletionWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetCompletionWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetCopyAction() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetEnableContinuousBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableContinuousBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetRecoveryPointTags() {
	_jsii_.InvokeVoid(
		a,
		"resetRecoveryPointTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetScanAction() {
	_jsii_.InvokeVoid(
		a,
		"resetScanAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetScheduleExpressionTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleExpressionTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetStartWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetStartWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ResetTargetLogicallyAirGappedBackupVaultArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetLogicallyAirGappedBackupVaultArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPlan_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

