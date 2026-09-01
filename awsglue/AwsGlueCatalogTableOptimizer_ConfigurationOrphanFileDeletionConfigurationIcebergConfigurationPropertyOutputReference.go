package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference interface {
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

// The jsii proxy struct for AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference
type jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) OrphanFileRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orphanFileRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) OrphanFileRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orphanFileRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) RunRateInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runRateInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference_Override(a AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetOrphanFileRetentionPeriodInDays(val *float64) {
	if err := j.validateSetOrphanFileRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orphanFileRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetRunRateInHours(val *float64) {
	if err := j.validateSetRunRateInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runRateInHours",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetOrphanFileRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetOrphanFileRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ResetRunRateInHours() {
	_jsii_.InvokeVoid(
		a,
		"resetRunRateInHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

