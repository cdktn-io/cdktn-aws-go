package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbackup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbackup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference interface {
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

// The jsii proxy struct for TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference
type jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ExcludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ExcludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) IncludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) IncludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) RecoveryPointTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) RecoveryPointTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) SelectionWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) SelectionWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.TfRestoreTestingPlan.RecoveryPointSelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference_Override(t TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.TfRestoreTestingPlan.RecoveryPointSelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetExcludeVaults(val *[]*string) {
	if err := j.validateSetExcludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeVaults",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetIncludeVaults(val *[]*string) {
	if err := j.validateSetIncludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVaults",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetRecoveryPointTypes(val *[]*string) {
	if err := j.validateSetRecoveryPointTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTypes",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetSelectionWindowDays(val *float64) {
	if err := j.validateSetSelectionWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionWindowDays",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ResetExcludeVaults() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludeVaults",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ResetSelectionWindowDays() {
	_jsii_.InvokeVoid(
		t,
		"resetSelectionWindowDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRestoreTestingPlan_RecoveryPointSelectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

