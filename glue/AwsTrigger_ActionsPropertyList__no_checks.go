//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTrigger_ActionsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTrigger_ActionsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTrigger_ActionsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTrigger_ActionsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTrigger_ActionsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTrigger_ActionsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTrigger_ActionsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTrigger_ActionsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

