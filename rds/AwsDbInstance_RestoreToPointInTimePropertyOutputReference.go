package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/rds/jsii"

	"github.com/cdktn-io/cdktn-aws-go/rds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDbInstance_RestoreToPointInTimePropertyOutputReference interface {
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
	InternalValue() *AwsDbInstance_RestoreToPointInTimeProperty
	// Experimental.
	SetInternalValue(val *AwsDbInstance_RestoreToPointInTimeProperty)
	// Experimental.
	RestoreTime() *string
	// Experimental.
	SetRestoreTime(val *string)
	// Experimental.
	RestoreTimeInput() *string
	// Experimental.
	SourceDbInstanceAutomatedBackupsArn() *string
	// Experimental.
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	// Experimental.
	SourceDbInstanceAutomatedBackupsArnInput() *string
	// Experimental.
	SourceDbInstanceIdentifier() *string
	// Experimental.
	SetSourceDbInstanceIdentifier(val *string)
	// Experimental.
	SourceDbInstanceIdentifierInput() *string
	// Experimental.
	SourceDbiResourceId() *string
	// Experimental.
	SetSourceDbiResourceId(val *string)
	// Experimental.
	SourceDbiResourceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseLatestRestorableTime() interface{}
	// Experimental.
	SetUseLatestRestorableTime(val interface{})
	// Experimental.
	UseLatestRestorableTimeInput() interface{}
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
	ResetRestoreTime()
	// Experimental.
	ResetSourceDbInstanceAutomatedBackupsArn()
	// Experimental.
	ResetSourceDbInstanceIdentifier()
	// Experimental.
	ResetSourceDbiResourceId()
	// Experimental.
	ResetUseLatestRestorableTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDbInstance_RestoreToPointInTimePropertyOutputReference
type jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) InternalValue() *AwsDbInstance_RestoreToPointInTimeProperty {
	var returns *AwsDbInstance_RestoreToPointInTimeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) RestoreTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbInstanceAutomatedBackupsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) SourceDbiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDbInstance_RestoreToPointInTimePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDbInstance_RestoreToPointInTimePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDbInstance_RestoreToPointInTimePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbInstance.RestoreToPointInTimePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDbInstance_RestoreToPointInTimePropertyOutputReference_Override(a AwsDbInstance_RestoreToPointInTimePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbInstance.RestoreToPointInTimePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetInternalValue(val *AwsDbInstance_RestoreToPointInTimeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetRestoreTime(val *string) {
	if err := j.validateSetRestoreTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreTime",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetSourceDbInstanceAutomatedBackupsArn(val *string) {
	if err := j.validateSetSourceDbInstanceAutomatedBackupsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetSourceDbInstanceIdentifier(val *string) {
	if err := j.validateSetSourceDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetSourceDbiResourceId(val *string) {
	if err := j.validateSetSourceDbiResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbiResourceId",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ResetRestoreTime() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ResetSourceDbInstanceAutomatedBackupsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDbInstanceAutomatedBackupsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ResetSourceDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ResetSourceDbiResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDbiResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		a,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDbInstance_RestoreToPointInTimePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

