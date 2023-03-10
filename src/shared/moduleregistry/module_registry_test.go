package moduleregistry

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleRegistry(t *testing.T) {
	t.Run("successs, struct param", func(t *testing.T) {

		act := func(ctx context.Context, param ReqModel) (ResModel, error) {
			res := ResModel{
				Data: param.Param + " act result",
			}
			return res, nil
		}

		moduleRegistry := NewModuleRegistry()

		RegisterSharedModule(moduleRegistry, "/act", act)
		res, err := moduleRegistry.CallModule("/act", context.Background(), ReqClientModel{Param: "paramdata"})
		assert.Nil(t, err)
		assert.NotNil(t, res)

		// parse result
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		var result ResClientModel
		err = json.Unmarshal(resBytes, &result)
		assert.Nil(t, err)
		assert.Equal(t, "paramdata act result", result.Data)
	})

	t.Run("successs, struct param & pointer result", func(t *testing.T) {

		act := func(ctx context.Context, param ReqModel) (*ResModel, error) {
			res := ResModel{
				Data: param.Param + " act result",
			}
			return &res, nil
		}

		moduleRegistry := NewModuleRegistry()

		RegisterSharedModule(moduleRegistry, "/act", act)
		res, err := moduleRegistry.CallModule("/act", context.Background(), ReqClientModel{Param: "paramdata"})
		assert.Nil(t, err)
		assert.NotNil(t, res)

		// parse result
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		var result ResClientModel
		err = json.Unmarshal(resBytes, &result)
		assert.Nil(t, err)
		assert.Equal(t, "paramdata act result", result.Data)
	})

	t.Run("successs, primitive type param & pointer result", func(t *testing.T) {

		act := func(ctx context.Context, param string) (string, error) {
			return param + " act result", nil
		}

		moduleRegistry := NewModuleRegistry()

		RegisterSharedModule(moduleRegistry, "/act", act)
		res, err := moduleRegistry.CallModule("/act", context.Background(), "paramdata")
		assert.Nil(t, err)
		assert.NotNil(t, res)

		result, ok := res.(string)
		assert.Equal(t, true, ok)
		assert.Equal(t, "paramdata act result", result)
	})

	t.Run("successs, primitive type param", func(t *testing.T) {

		act := func(ctx context.Context, param string) (*string, error) {
			res := param + " act result"
			return &res, nil
		}

		moduleRegistry := NewModuleRegistry()

		RegisterSharedModule(moduleRegistry, "/act", act)
		res, err := moduleRegistry.CallModule("/act", context.Background(), "paramdata")
		assert.Nil(t, err)
		assert.NotNil(t, res)

		result, ok := res.(*string)
		assert.Equal(t, true, ok)
		assert.Equal(t, "paramdata act result", *result)
	})

	t.Run("failed, invalid param", func(t *testing.T) {

		act := func(ctx context.Context, param ReqModel) (ResModel, error) {
			res := ResModel{
				Data: param.Param + " act result",
			}
			return res, nil
		}

		moduleRegistry := NewModuleRegistry()

		RegisterSharedModule(moduleRegistry, "/act", act)
		res, err := moduleRegistry.CallModule("/act", context.Background(), "paramdata")
		assert.NotNil(t, err)
		err, ok := err.(InvalidParameter)
		assert.Equal(t, true, ok)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

type ReqModel struct {
	Param string
}

type ReqClientModel struct {
	Param string
}

type ResModel struct {
	Data string
}

type ResClientModel struct {
	Data string
}
