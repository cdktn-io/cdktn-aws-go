package backup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/backup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/backup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Algorithm() *string
	// Experimental.
	SetAlgorithm(val *string)
	// Experimental.
	AlgorithmInput() *string
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
	ExcludeVaults() *[]*string
	// Experimental.
	SetExcludeVaults(val *[]*string)
	// Experimental.
	ExcludeVaultsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeVaults() *[]*string
	// Experimental.
	SetIncludeVaults(val *[]*string)
	// Experimental.
	IncludeVaultsInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RecoveryPointTypes() *[]*string
	// Experimental.
	SetRecoveryPointTypes(val *[]*string)
	// Experimental.
	RecoveryPointTypesInput() *[]*string
	// Experimental.
	SelectionWindowDays() *float64
	// Experimental.
	SetSelectionWindowDays(val *float64)
	// Experimental.
	SelectionWindowDaysInput() *float64
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
	ResetExcludeVaults()
	// Experimental.
	ResetSelectionWindowDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference
type jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ExcludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ExcludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) IncludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) IncludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) RecoveryPointTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) RecoveryPointTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) SelectionWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) SelectionWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.AwsRestoreTestingPlan.RecoveryPointSelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference_Override(a AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.AwsRestoreTestingPlan.RecoveryPointSelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetExcludeVaults(val *[]*string) {
	if err := j.validateSetExcludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeVaults",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetIncludeVaults(val *[]*string) {
	if err := j.validateSetIncludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVaults",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetRecoveryPointTypes(val *[]*string) {
	if err := j.validateSetRecoveryPointTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTypes",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetSelectionWindowDays(val *float64) {
	if err := j.validateSetSelectionWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionWindowDays",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ResetExcludeVaults() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeVaults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ResetSelectionWindowDays() {
	_jsii_.InvokeVoid(
		a,
		"resetSelectionWindowDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

