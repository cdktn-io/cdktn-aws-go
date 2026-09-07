package kinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_ApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApplicationCodeConfiguration() AwsApplication_ApplicationCodeConfigurationPropertyOutputReference
	// Experimental.
	ApplicationCodeConfigurationInput() *AwsApplication_ApplicationCodeConfigurationProperty
	// Experimental.
	ApplicationEncryptionConfiguration() AwsApplication_ApplicationEncryptionConfigurationPropertyOutputReference
	// Experimental.
	ApplicationEncryptionConfigurationInput() *AwsApplication_ApplicationEncryptionConfigurationProperty
	// Experimental.
	ApplicationSnapshotConfiguration() AwsApplication_ApplicationSnapshotConfigurationPropertyOutputReference
	// Experimental.
	ApplicationSnapshotConfigurationInput() *AwsApplication_ApplicationSnapshotConfigurationProperty
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
	EnvironmentProperties() AwsApplication_EnvironmentPropertiesPropertyOutputReference
	// Experimental.
	EnvironmentPropertiesInput() *AwsApplication_EnvironmentPropertiesProperty
	// Experimental.
	FlinkApplicationConfiguration() AwsApplication_FlinkApplicationConfigurationPropertyOutputReference
	// Experimental.
	FlinkApplicationConfigurationInput() *AwsApplication_FlinkApplicationConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsApplication_ApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_ApplicationConfigurationProperty)
	// Experimental.
	RunConfiguration() AwsApplication_RunConfigurationPropertyOutputReference
	// Experimental.
	RunConfigurationInput() *AwsApplication_RunConfigurationProperty
	// Experimental.
	SqlApplicationConfiguration() AwsApplication_SqlApplicationConfigurationPropertyOutputReference
	// Experimental.
	SqlApplicationConfigurationInput() *AwsApplication_SqlApplicationConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcConfiguration() AwsApplication_VpcConfigurationPropertyOutputReference
	// Experimental.
	VpcConfigurationInput() *AwsApplication_VpcConfigurationProperty
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
	PutApplicationCodeConfiguration(value *AwsApplication_ApplicationCodeConfigurationProperty)
	// Experimental.
	PutApplicationEncryptionConfiguration(value *AwsApplication_ApplicationEncryptionConfigurationProperty)
	// Experimental.
	PutApplicationSnapshotConfiguration(value *AwsApplication_ApplicationSnapshotConfigurationProperty)
	// Experimental.
	PutEnvironmentProperties(value *AwsApplication_EnvironmentPropertiesProperty)
	// Experimental.
	PutFlinkApplicationConfiguration(value *AwsApplication_FlinkApplicationConfigurationProperty)
	// Experimental.
	PutRunConfiguration(value *AwsApplication_RunConfigurationProperty)
	// Experimental.
	PutSqlApplicationConfiguration(value *AwsApplication_SqlApplicationConfigurationProperty)
	// Experimental.
	PutVpcConfiguration(value *AwsApplication_VpcConfigurationProperty)
	// Experimental.
	ResetApplicationEncryptionConfiguration()
	// Experimental.
	ResetApplicationSnapshotConfiguration()
	// Experimental.
	ResetEnvironmentProperties()
	// Experimental.
	ResetFlinkApplicationConfiguration()
	// Experimental.
	ResetRunConfiguration()
	// Experimental.
	ResetSqlApplicationConfiguration()
	// Experimental.
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApplication_ApplicationConfigurationPropertyOutputReference
type jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfiguration() AwsApplication_ApplicationCodeConfigurationPropertyOutputReference {
	var returns AwsApplication_ApplicationCodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationCodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfigurationInput() *AwsApplication_ApplicationCodeConfigurationProperty {
	var returns *AwsApplication_ApplicationCodeConfigurationProperty
	_jsii_.Get(
		j,
		"applicationCodeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfiguration() AwsApplication_ApplicationEncryptionConfigurationPropertyOutputReference {
	var returns AwsApplication_ApplicationEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfigurationInput() *AwsApplication_ApplicationEncryptionConfigurationProperty {
	var returns *AwsApplication_ApplicationEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"applicationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfiguration() AwsApplication_ApplicationSnapshotConfigurationPropertyOutputReference {
	var returns AwsApplication_ApplicationSnapshotConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationSnapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfigurationInput() *AwsApplication_ApplicationSnapshotConfigurationProperty {
	var returns *AwsApplication_ApplicationSnapshotConfigurationProperty
	_jsii_.Get(
		j,
		"applicationSnapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) EnvironmentProperties() AwsApplication_EnvironmentPropertiesPropertyOutputReference {
	var returns AwsApplication_EnvironmentPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"environmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) EnvironmentPropertiesInput() *AwsApplication_EnvironmentPropertiesProperty {
	var returns *AwsApplication_EnvironmentPropertiesProperty
	_jsii_.Get(
		j,
		"environmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfiguration() AwsApplication_FlinkApplicationConfigurationPropertyOutputReference {
	var returns AwsApplication_FlinkApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"flinkApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfigurationInput() *AwsApplication_FlinkApplicationConfigurationProperty {
	var returns *AwsApplication_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"flinkApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) InternalValue() *AwsApplication_ApplicationConfigurationProperty {
	var returns *AwsApplication_ApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) RunConfiguration() AwsApplication_RunConfigurationPropertyOutputReference {
	var returns AwsApplication_RunConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) RunConfigurationInput() *AwsApplication_RunConfigurationProperty {
	var returns *AwsApplication_RunConfigurationProperty
	_jsii_.Get(
		j,
		"runConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfiguration() AwsApplication_SqlApplicationConfigurationPropertyOutputReference {
	var returns AwsApplication_SqlApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfigurationInput() *AwsApplication_SqlApplicationConfigurationProperty {
	var returns *AwsApplication_SqlApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"sqlApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) VpcConfiguration() AwsApplication_VpcConfigurationPropertyOutputReference {
	var returns AwsApplication_VpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) VpcConfigurationInput() *AwsApplication_VpcConfigurationProperty {
	var returns *AwsApplication_VpcConfigurationProperty
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_ApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_ApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_ApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_ApplicationConfigurationPropertyOutputReference_Override(a AwsApplication_ApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference)SetInternalValue(val *AwsApplication_ApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationCodeConfiguration(value *AwsApplication_ApplicationCodeConfigurationProperty) {
	if err := a.validatePutApplicationCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationEncryptionConfiguration(value *AwsApplication_ApplicationEncryptionConfigurationProperty) {
	if err := a.validatePutApplicationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationSnapshotConfiguration(value *AwsApplication_ApplicationSnapshotConfigurationProperty) {
	if err := a.validatePutApplicationSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutEnvironmentProperties(value *AwsApplication_EnvironmentPropertiesProperty) {
	if err := a.validatePutEnvironmentPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutFlinkApplicationConfiguration(value *AwsApplication_FlinkApplicationConfigurationProperty) {
	if err := a.validatePutFlinkApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFlinkApplicationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutRunConfiguration(value *AwsApplication_RunConfigurationProperty) {
	if err := a.validatePutRunConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutSqlApplicationConfiguration(value *AwsApplication_SqlApplicationConfigurationProperty) {
	if err := a.validatePutSqlApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlApplicationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) PutVpcConfiguration(value *AwsApplication_VpcConfigurationProperty) {
	if err := a.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetApplicationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetApplicationSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSnapshotConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetEnvironmentProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetFlinkApplicationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFlinkApplicationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetRunConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRunConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetSqlApplicationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlApplicationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_ApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

