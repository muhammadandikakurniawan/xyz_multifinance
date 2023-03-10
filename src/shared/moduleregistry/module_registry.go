package moduleregistry

import (
	"context"
	"encoding/json"
	"fmt"
)

func NewModuleRegistry() ModuleRegistry {
	return &moduleRegistryImpl{
		modules: map[string]func(ctx context.Context, param any) (any, error){},
	}
}

type ModuleRegistry interface {
	RegisterModule(path string, actModule func(ctx context.Context, param any) (any, error))
	CallModule(path string, ctx context.Context, param any) (res any, err error)
}

type moduleRegistryImpl struct {
	modules map[string]func(ctx context.Context, param any) (any, error)
}

func (m *moduleRegistryImpl) RegisterModule(path string, actModule func(ctx context.Context, param any) (any, error)) {
	m.modules[path] = actModule
}

func (m *moduleRegistryImpl) CallModule(path string, ctx context.Context, param any) (res any, err error) {
	act, exists := m.modules[path]
	if !exists {
		err = fmt.Errorf("module not found")
		return
	}

	return act(ctx, param)
}

func RegisterSharedModule[T_PARAM any, T_RESULT any](registry ModuleRegistry, path string, act func(ctx context.Context, param T_PARAM) (T_RESULT, error)) {

	regAct := func(ctx context.Context, param any) (result any, err error) {
		paramParse, valid := param.(T_PARAM)
		if !valid {

			paramBytes, jsonErr := json.Marshal(param)
			if jsonErr != nil {
				err = NewInvalidParameter(jsonErr.Error())
				return
			}

			if jsonErr = json.Unmarshal(paramBytes, &paramParse); jsonErr != nil {
				err = NewInvalidParameter(jsonErr.Error())
				return
			}
		}

		result, err = act(ctx, paramParse)
		return
	}
	registry.RegisterModule(path, regAct)
}
