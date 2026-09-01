package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference interface {
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
	DeleteFileThreshold() *float64
	// Experimental.
	SetDeleteFileThreshold(val *float64)
	// Experimental.
	DeleteFileThresholdInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MinInputFiles() *float64
	// Experimental.
	SetMinInputFiles(val *float64)
	// Experimental.
	MinInputFilesInput() *float64
	// Experimental.
	Strategy() *string
	// Experimental.
	SetStrategy(val *string)
	// Experimental.
	StrategyInput() *string
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
	ResetDeleteFileThreshold()
	// Experimental.
	ResetMinInputFiles()
	// Experimental.
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference
type jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) DeleteFileThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteFileThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) DeleteFileThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteFileThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) MinInputFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInputFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) MinInputFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInputFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTableOptimizer.ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference_Override(a AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTableOptimizer.ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetDeleteFileThreshold(val *float64) {
	if err := j.validateSetDeleteFileThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteFileThreshold",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetMinInputFiles(val *float64) {
	if err := j.validateSetMinInputFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInputFiles",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ResetDeleteFileThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteFileThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ResetMinInputFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetMinInputFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

