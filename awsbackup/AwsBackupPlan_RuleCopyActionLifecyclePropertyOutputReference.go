package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbackup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbackup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ColdStorageAfter() *float64
	// Experimental.
	SetColdStorageAfter(val *float64)
	// Experimental.
	ColdStorageAfterInput() *float64
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
	DeleteAfter() *float64
	// Experimental.
	SetDeleteAfter(val *float64)
	// Experimental.
	DeleteAfterInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsBackupPlan_RuleCopyActionLifecycleProperty
	// Experimental.
	SetInternalValue(val *AwsBackupPlan_RuleCopyActionLifecycleProperty)
	// Experimental.
	OptInToArchiveForSupportedResources() interface{}
	// Experimental.
	SetOptInToArchiveForSupportedResources(val interface{})
	// Experimental.
	OptInToArchiveForSupportedResourcesInput() interface{}
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
	ResetColdStorageAfter()
	// Experimental.
	ResetDeleteAfter()
	// Experimental.
	ResetOptInToArchiveForSupportedResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference
type jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ColdStorageAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coldStorageAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ColdStorageAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coldStorageAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) DeleteAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) DeleteAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) InternalValue() *AwsBackupPlan_RuleCopyActionLifecycleProperty {
	var returns *AwsBackupPlan_RuleCopyActionLifecycleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) OptInToArchiveForSupportedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) OptInToArchiveForSupportedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupPlan.RuleCopyActionLifecyclePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference_Override(a AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupPlan.RuleCopyActionLifecyclePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetColdStorageAfter(val *float64) {
	if err := j.validateSetColdStorageAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coldStorageAfter",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetDeleteAfter(val *float64) {
	if err := j.validateSetDeleteAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfter",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetInternalValue(val *AwsBackupPlan_RuleCopyActionLifecycleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetOptInToArchiveForSupportedResources(val interface{}) {
	if err := j.validateSetOptInToArchiveForSupportedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optInToArchiveForSupportedResources",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ResetColdStorageAfter() {
	_jsii_.InvokeVoid(
		a,
		"resetColdStorageAfter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ResetDeleteAfter() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAfter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ResetOptInToArchiveForSupportedResources() {
	_jsii_.InvokeVoid(
		a,
		"resetOptInToArchiveForSupportedResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBackupPlan_RuleCopyActionLifecyclePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

