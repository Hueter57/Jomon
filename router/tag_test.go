package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/traPtitech/Jomon/model"
	"github.com/traPtitech/Jomon/testutil/random"
)

func TestHandlers_GetTags(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)
		date := time.Now()

		tag1 := &model.Tag{
			ID:          uuid.New(),
			Name:        random.AlphaNumeric(t, 20),
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}
		tag2 := &model.Tag{
			ID:          uuid.New(),
			Name:        random.AlphaNumeric(t, 20),
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}
		tags := []*model.Tag{tag1, tag2}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			GetTags(ctx).
			Return(tags, nil)

		var resBody TagResponse
		statusCode, _ := th.doRequest(t, echo.GET, "/api/tags", nil, &resBody)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Len(t, resBody.Tags, 2)
		if resBody.Tags[0].ID == tag1.ID {
			assert.Equal(t, resBody.Tags[0].ID, tag1.ID)
			assert.Equal(t, resBody.Tags[0].Name, tag1.Name)
			assert.Equal(t, resBody.Tags[0].Description, tag1.Description)
			assert.Equal(t, resBody.Tags[1].ID, tag2.ID)
			assert.Equal(t, resBody.Tags[1].Name, tag2.Name)
			assert.Equal(t, resBody.Tags[1].Description, tag2.Description)
		} else {
			assert.Equal(t, resBody.Tags[0].ID, tag2.ID)
			assert.Equal(t, resBody.Tags[0].Name, tag2.Name)
			assert.Equal(t, resBody.Tags[0].Description, tag2.Description)
			assert.Equal(t, resBody.Tags[1].ID, tag1.ID)
			assert.Equal(t, resBody.Tags[1].Name, tag1.Name)
			assert.Equal(t, resBody.Tags[1].Description, tag1.Description)
		}
	})

	t.Run("Success2", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)

		tags := []*model.Tag{}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			GetTags(ctx).
			Return(tags, nil)

		var resBody TagResponse
		statusCode, _ := th.doRequest(t, echo.GET, "/api/tags", nil, &resBody)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Len(t, resBody.Tags, 0)
	})
}

func TestHandlers_PostTag(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)
		date := time.Now()

		tag := &model.Tag{
			ID:          uuid.New(),
			Name:        random.AlphaNumeric(t, 20),
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			CreateTag(ctx, tag.Name, tag.Description).
			Return(tag, nil)

		req := Tag{
			Name:        tag.Name,
			Description: tag.Description,
		}

		var resBody TagOverview
		statusCode, _ := th.doRequest(t, echo.POST, "/api/tags", &req, &resBody)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Equal(t, tag.ID, resBody.ID)
		assert.Equal(t, tag.Name, resBody.Name)
		assert.Equal(t, tag.Description, resBody.Description)
	})

	t.Run("MissingName", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)
		date := time.Now()

		tag := &model.Tag{
			ID:          uuid.New(),
			Name:        "",
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			CreateTag(ctx, tag.Name, tag.Description).
			Return(nil, errors.New("Tag name can't empty."))

		req := Tag{
			Name:        tag.Name,
			Description: tag.Description,
		}

		var resBody TagOverview
		statusCode, _ := th.doRequest(t, echo.POST, "/api/tags", &req, &resBody)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})
}

func TestHandlers_PutTag(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)
		date := time.Now()

		tag := &model.Tag{
			ID:          uuid.New(),
			Name:        random.AlphaNumeric(t, 20),
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			UpdateTag(ctx, tag.ID, tag.Name, tag.Description).
			Return(tag, nil)

		req := Tag{
			Name:        tag.Name,
			Description: tag.Description,
		}

		var resBody TagOverview
		path := fmt.Sprintf("/api/tags/%s", tag.ID.String())
		statusCode, _ := th.doRequest(t, echo.PUT, path, &req, &resBody)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Equal(t, tag.ID, resBody.ID)
		assert.Equal(t, tag.Name, resBody.Name)
		assert.Equal(t, tag.Description, resBody.Description)
	})

	t.Run("MissingName", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		th, err := SetupTestHandlers(t, ctrl)
		assert.NoError(t, err)
		date := time.Now()

		tag := &model.Tag{
			ID:          uuid.New(),
			Name:        "",
			Description: random.AlphaNumeric(t, 50),
			CreatedAt:   date,
			UpdatedAt:   date,
		}

		ctx := context.Background()
		th.Repository.MockTagRepository.
			EXPECT().
			UpdateTag(ctx, tag.ID, tag.Name, tag.Description).
			Return(nil, errors.New("Tag name can't empty."))

		req := Tag{
			Name:        tag.Name,
			Description: tag.Description,
		}

		var resBody TagOverview
		path := fmt.Sprintf("/api/tags/%s", tag.ID.String())
		statusCode, _ := th.doRequest(t, echo.PUT, path, &req, &resBody)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})
}
