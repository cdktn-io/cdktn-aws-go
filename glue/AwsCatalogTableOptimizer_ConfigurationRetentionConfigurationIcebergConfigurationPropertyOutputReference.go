package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CleanExpiredFiles() interface{}
	// Experimental.
	SetCleanExpiredFiles(val interface{})
	// Experimental.
	CleanExpiredFilesInput() interface{}
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
	NumberOfSnapshotsToRetain() *float64
	// Experimental.
	SetNumberOfSnapshotsToRetain(val *float64)
	// Experimental.
	NumberOfSnapshotsToRetainInput() *float64
	// Experimental.
	RunRateInHours() *float64
	// Experimental.
	SetRunRateInHours(val *float64)
	// Experimental.
	RunRateInHoursInput() *float64
	// Experimental.
	SnapshotRetentionPeriodInDays() *float64
	// Experimental.
	SetSnapshotRetentionPeriodInDays(val *float64)
	// Experimental.
	SnapshotRetentionPeriodInDaysInput() *float64
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
	ResetCleanExpiredFiles()
	// Experimental.
	ResetNumberOfSnapshotsToRetain()
	// Experimental.
	ResetRunRateInHours()
	// Experimental.
	ResetSnapshotRetentionPeriodInDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference
type jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) CleanExpiredFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanExpiredFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) CleanExpiredFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanExpiredFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) NumberOfSnapshotsToRetain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSnapshotsToRetain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) NumberOfSnapshotsToRetainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSnapshotsToRetainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) SnapshotRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) SnapshotRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTableOptimizer.ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference_Override(a AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTableOptimizer.ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetCleanExpiredFiles(val interface{}) {
	if err := j.validateSetCleanExpiredFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanExpiredFiles",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetNumberOfSnapshotsToRetain(val *float64) {
	if err := j.validateSetNumberOfSnapshotsToRetainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfSnapshotsToRetain",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetRunRateInHours(val *float64) {
	if err := j.validateSetRunRateInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runRateInHours",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetSnapshotRetentionPeriodInDays(val *float64) {
	if err := j.validateSetSnapshotRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ResetCleanExpiredFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetCleanExpiredFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ResetNumberOfSnapshotsToRetain() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfSnapshotsToRetain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ResetRunRateInHours() {
	_jsii_.InvokeVoid(
		a,
		"resetRunRateInHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ResetSnapshotRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

