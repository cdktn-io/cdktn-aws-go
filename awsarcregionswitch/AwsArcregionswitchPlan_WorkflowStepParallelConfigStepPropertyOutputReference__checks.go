//go:build !no_runtime_type_checking

package awsarcregionswitch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutArcRoutingControlConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutAuroraProvisionedScalingConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraProvisionedScalingConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutAuroraServerlessScalingConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepAuroraServerlessScalingConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutCustomActionLambdaConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutDocumentDbConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepDocumentDbConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepDocumentDbConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepDocumentDbConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepDocumentDbConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepDocumentDbConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutEc2AsgCapacityIncreaseConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutEcsCapacityIncreaseConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutEksResourceScalingConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutExecutionApprovalConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutGlobalAuroraConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepGlobalAuroraConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutLambdaEventSourceMappingConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutNeptuneGlobalDatabaseConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepNeptuneGlobalDatabaseConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutRdsCreateCrossRegionReadReplicaConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsCreateCrossRegionReadReplicaConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutRdsPromoteReadReplicaConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRdsPromoteReadReplicaConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutRegionSwitchPlanConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRegionSwitchPlanConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validatePutRoute53HealthCheckConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty:
		value := value.(*[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty:
		value_ := value.([]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetExecutionBlockTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty:
		val := val.(*AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty:
		val_ := val.(AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsArcregionswitchPlan_WorkflowStepParallelConfigStepPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

