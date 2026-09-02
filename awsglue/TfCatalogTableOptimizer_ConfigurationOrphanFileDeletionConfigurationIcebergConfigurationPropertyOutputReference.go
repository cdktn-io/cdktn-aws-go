package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference interface {
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
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	OrphanFileRetentionPeriodInDays() *float64
	// Experimental.
	SetOrphanFileRetentionPeriodInDays(val *float64)
	// Experimental.
	OrphanFileRetentionPeriodInDaysInput() *float64
	// Experimental.
	RunRateInHours() *float64
	// Experimental.
	SetRunRateInHours(val *float64)
	// Experimental.
	RunRateInHoursInput() *float64
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
	ResetLocation()
	// Experimental.
	ResetOrphanFileRetentionPeriodInDays()
	// Experimental.
	ResetRunRateInHours()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference
type jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) OrphanFileRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orphanFileRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) OrphanFileRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orphanFileRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference_Override(t TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetOrphanFileRetentionPeriodInDays(val *float64) {
	if err := j.validateSetOrphanFileRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orphanFileRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetRunRateInHours(val *float64) {
	if err := j.validateSetRunRateInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runRateInHours",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetOrphanFileRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		t,
		"resetOrphanFileRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetRunRateInHours() {
	_jsii_.InvokeVoid(
		t,
		"resetRunRateInHours",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

