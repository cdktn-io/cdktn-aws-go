package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_ApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApplicationCodeConfiguration() TfApplication_ApplicationCodeConfigurationPropertyOutputReference
	// Experimental.
	ApplicationCodeConfigurationInput() *TfApplication_ApplicationCodeConfigurationProperty
	// Experimental.
	ApplicationEncryptionConfiguration() TfApplication_ApplicationEncryptionConfigurationPropertyOutputReference
	// Experimental.
	ApplicationEncryptionConfigurationInput() *TfApplication_ApplicationEncryptionConfigurationProperty
	// Experimental.
	ApplicationSnapshotConfiguration() TfApplication_ApplicationSnapshotConfigurationPropertyOutputReference
	// Experimental.
	ApplicationSnapshotConfigurationInput() *TfApplication_ApplicationSnapshotConfigurationProperty
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
	EnvironmentProperties() TfApplication_EnvironmentPropertiesPropertyOutputReference
	// Experimental.
	EnvironmentPropertiesInput() *TfApplication_EnvironmentPropertiesProperty
	// Experimental.
	FlinkApplicationConfiguration() TfApplication_FlinkApplicationConfigurationPropertyOutputReference
	// Experimental.
	FlinkApplicationConfigurationInput() *TfApplication_FlinkApplicationConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfApplication_ApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfApplication_ApplicationConfigurationProperty)
	// Experimental.
	RunConfiguration() TfApplication_RunConfigurationPropertyOutputReference
	// Experimental.
	RunConfigurationInput() *TfApplication_RunConfigurationProperty
	// Experimental.
	SqlApplicationConfiguration() TfApplication_SqlApplicationConfigurationPropertyOutputReference
	// Experimental.
	SqlApplicationConfigurationInput() *TfApplication_SqlApplicationConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcConfiguration() TfApplication_VpcConfigurationPropertyOutputReference
	// Experimental.
	VpcConfigurationInput() *TfApplication_VpcConfigurationProperty
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
	PutApplicationCodeConfiguration(value *TfApplication_ApplicationCodeConfigurationProperty)
	// Experimental.
	PutApplicationEncryptionConfiguration(value *TfApplication_ApplicationEncryptionConfigurationProperty)
	// Experimental.
	PutApplicationSnapshotConfiguration(value *TfApplication_ApplicationSnapshotConfigurationProperty)
	// Experimental.
	PutEnvironmentProperties(value *TfApplication_EnvironmentPropertiesProperty)
	// Experimental.
	PutFlinkApplicationConfiguration(value *TfApplication_FlinkApplicationConfigurationProperty)
	// Experimental.
	PutRunConfiguration(value *TfApplication_RunConfigurationProperty)
	// Experimental.
	PutSqlApplicationConfiguration(value *TfApplication_SqlApplicationConfigurationProperty)
	// Experimental.
	PutVpcConfiguration(value *TfApplication_VpcConfigurationProperty)
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

// The jsii proxy struct for TfApplication_ApplicationConfigurationPropertyOutputReference
type jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfiguration() TfApplication_ApplicationCodeConfigurationPropertyOutputReference {
	var returns TfApplication_ApplicationCodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationCodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfigurationInput() *TfApplication_ApplicationCodeConfigurationProperty {
	var returns *TfApplication_ApplicationCodeConfigurationProperty
	_jsii_.Get(
		j,
		"applicationCodeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfiguration() TfApplication_ApplicationEncryptionConfigurationPropertyOutputReference {
	var returns TfApplication_ApplicationEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfigurationInput() *TfApplication_ApplicationEncryptionConfigurationProperty {
	var returns *TfApplication_ApplicationEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"applicationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfiguration() TfApplication_ApplicationSnapshotConfigurationPropertyOutputReference {
	var returns TfApplication_ApplicationSnapshotConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationSnapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfigurationInput() *TfApplication_ApplicationSnapshotConfigurationProperty {
	var returns *TfApplication_ApplicationSnapshotConfigurationProperty
	_jsii_.Get(
		j,
		"applicationSnapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) EnvironmentProperties() TfApplication_EnvironmentPropertiesPropertyOutputReference {
	var returns TfApplication_EnvironmentPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"environmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) EnvironmentPropertiesInput() *TfApplication_EnvironmentPropertiesProperty {
	var returns *TfApplication_EnvironmentPropertiesProperty
	_jsii_.Get(
		j,
		"environmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfiguration() TfApplication_FlinkApplicationConfigurationPropertyOutputReference {
	var returns TfApplication_FlinkApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"flinkApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfigurationInput() *TfApplication_FlinkApplicationConfigurationProperty {
	var returns *TfApplication_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"flinkApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) InternalValue() *TfApplication_ApplicationConfigurationProperty {
	var returns *TfApplication_ApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) RunConfiguration() TfApplication_RunConfigurationPropertyOutputReference {
	var returns TfApplication_RunConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) RunConfigurationInput() *TfApplication_RunConfigurationProperty {
	var returns *TfApplication_RunConfigurationProperty
	_jsii_.Get(
		j,
		"runConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfiguration() TfApplication_SqlApplicationConfigurationPropertyOutputReference {
	var returns TfApplication_SqlApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfigurationInput() *TfApplication_SqlApplicationConfigurationProperty {
	var returns *TfApplication_SqlApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"sqlApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) VpcConfiguration() TfApplication_VpcConfigurationPropertyOutputReference {
	var returns TfApplication_VpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) VpcConfigurationInput() *TfApplication_VpcConfigurationProperty {
	var returns *TfApplication_VpcConfigurationProperty
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_ApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_ApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_ApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_ApplicationConfigurationPropertyOutputReference_Override(t TfApplication_ApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference)SetInternalValue(val *TfApplication_ApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationCodeConfiguration(value *TfApplication_ApplicationCodeConfigurationProperty) {
	if err := t.validatePutApplicationCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApplicationCodeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationEncryptionConfiguration(value *TfApplication_ApplicationEncryptionConfigurationProperty) {
	if err := t.validatePutApplicationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApplicationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutApplicationSnapshotConfiguration(value *TfApplication_ApplicationSnapshotConfigurationProperty) {
	if err := t.validatePutApplicationSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApplicationSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutEnvironmentProperties(value *TfApplication_EnvironmentPropertiesProperty) {
	if err := t.validatePutEnvironmentPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironmentProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutFlinkApplicationConfiguration(value *TfApplication_FlinkApplicationConfigurationProperty) {
	if err := t.validatePutFlinkApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFlinkApplicationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutRunConfiguration(value *TfApplication_RunConfigurationProperty) {
	if err := t.validatePutRunConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRunConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutSqlApplicationConfiguration(value *TfApplication_SqlApplicationConfigurationProperty) {
	if err := t.validatePutSqlApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqlApplicationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) PutVpcConfiguration(value *TfApplication_VpcConfigurationProperty) {
	if err := t.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetApplicationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetApplicationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetApplicationSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetApplicationSnapshotConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetEnvironmentProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetFlinkApplicationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetFlinkApplicationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetRunConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetRunConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetSqlApplicationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSqlApplicationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

