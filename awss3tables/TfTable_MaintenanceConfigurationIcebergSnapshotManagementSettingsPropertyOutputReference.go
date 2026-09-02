package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3tables/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3tables/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxSnapshotAgeHours() *float64
	// Experimental.
	SetMaxSnapshotAgeHours(val *float64)
	// Experimental.
	MaxSnapshotAgeHoursInput() *float64
	// Experimental.
	MinSnapshotsToKeep() *float64
	// Experimental.
	SetMinSnapshotsToKeep(val *float64)
	// Experimental.
	MinSnapshotsToKeepInput() *float64
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
	ResetMaxSnapshotAgeHours()
	// Experimental.
	ResetMinSnapshotsToKeep()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference
type jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) MaxSnapshotAgeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSnapshotAgeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) MaxSnapshotAgeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSnapshotAgeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) MinSnapshotsToKeep() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSnapshotsToKeep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) MinSnapshotsToKeepInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSnapshotsToKeepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-tables.TfTable.MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference_Override(t TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-tables.TfTable.MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetMaxSnapshotAgeHours(val *float64) {
	if err := j.validateSetMaxSnapshotAgeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSnapshotAgeHours",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetMinSnapshotsToKeep(val *float64) {
	if err := j.validateSetMinSnapshotsToKeepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSnapshotsToKeep",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ResetMaxSnapshotAgeHours() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxSnapshotAgeHours",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ResetMinSnapshotsToKeep() {
	_jsii_.InvokeVoid(
		t,
		"resetMinSnapshotsToKeep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

