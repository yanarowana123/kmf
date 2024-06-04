package http

import (
	"context"
	"encoding/json"
	_ "github.com/yanarowana123/kmf/cmd/docs"
	"github.com/yanarowana123/kmf/internal/model"
	"net/http"
	"time"
)

type CurrencyFetcherService interface {
	Save(ctx context.Context, date time.Time) error
	List(ctx context.Context, date time.Time, code string) ([]model.GetCurrencyResponse, error)
}

type Controller struct {
	s CurrencyFetcherService
}

func NewController(s CurrencyFetcherService) Controller {
	return Controller{s}
}

// Save
// @Summary save currency
// @Description save currency
// @Success 200 {object} model.SaveCurrencyResponse
// @Failure 400 "Bad request"
// @Router /currency/save/{date} [post]
func (c Controller) Save(ctx context.Context, w http.ResponseWriter, r *http.Request, dateString string) {
	date, err := time.Parse("02.01.2006", dateString)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "invalid date format. Correct format is 02.01.2006(day.month.year)")
		return
	}

	err = c.s.Save(ctx, date)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.SaveCurrencyResponse{Success: true})
}

func (c Controller) List(ctx context.Context, w http.ResponseWriter, r *http.Request, dateString, code string) {
	date, err := time.Parse("02.01.2006", dateString)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "invalid date format. Correct format is 02.01.2006(day.month.year) EEE")
		return
	}

	resp, err := c.s.List(ctx, date, code)

	if err != nil {
		c.respondWithError(w, http.StatusInternalServerError, "unexpected error")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

//// Create
//// @Summary create task
//// @Description creates task
//// @Param task body dto.UpsertTaskSwagger true "body"
//// @Success 200 {object} dto.CreateTaskResponse
//// @Failure 404 "Something went wrong while creating"
//// @Failure 400 "Bad request"
//// @Router /api/todo-list/tasks [post]
//func (c Controller) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
//	var createTaskRequest dto.UpsertTaskRequest
//	dec := json.NewDecoder(r.Body)
//	dec.DisallowUnknownFields()
//
//	err := dec.Decode(&createTaskRequest)
//	if err != nil {
//		c.respondWithError(w, http.StatusBadRequest, "Bad request")
//		return
//	}
//
//	// маппинг сущности из одного слоя в другой можно вынести в отдельный пакет
//	id, err := c.s.Create(ctx, applicationDto.UpsertTask{Title: createTaskRequest.Title, ActiveAt: createTaskRequest.ActiveAt.Time})
//
//	if err != nil {
//		c.respondWithError(w, http.StatusNotFound, "")
//		return
//	}
//
//	json.NewEncoder(w).Encode(dto.CreateTaskResponse{Id: id})
//}
//
//// Update принимает id извне так как для его получения используется специфичный метод роутера (сегодня chi, завтра echo)
//// @Summary update task
//// @Description updates task
//// @Param id path string true "task id"
//// @Param task body dto.UpsertTaskSwagger true "body"
//// @Success 204
//// @Failure 404 "Task not found"
//// @Failure 400 "Bad request"
//// @Router /api/todo-list/tasks/{id} [put]
//func (c Controller) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) {
//	var updateTaskRequest dto.UpsertTaskRequest
//	dec := json.NewDecoder(r.Body)
//	dec.DisallowUnknownFields()
//
//	err := dec.Decode(&updateTaskRequest)
//	if err != nil {
//		c.respondWithError(w, http.StatusBadRequest, "Bad request")
//		return
//	}
//
//	err = c.s.Update(ctx, id, applicationDto.UpsertTask{Title: updateTaskRequest.Title, ActiveAt: updateTaskRequest.ActiveAt.Time})
//
//	if err != nil {
//		c.respondWithError(w, http.StatusNotFound, "")
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
//
//// Delete принимает id извне так как для его получения используется специфичный метод роутера (сегодня chi, завтра echo)
//// @Summary delete task
//// @Description deletes task
//// @Param id path string true "task id"
//// @Success 204
//// @Failure 404 "Task not found"
//// @Router /api/todo-list/tasks/{id} [delete]
//func (c Controller) Delete(ctx context.Context, w http.ResponseWriter, id string) {
//	err := c.s.Delete(ctx, id)
//
//	if err != nil {
//		c.respondWithError(w, http.StatusNotFound, err.Error())
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
//
//// Done принимает id извне так как для его получения используется специфичный метод роутера (сегодня chi, завтра echo)
//// @Summary done task
//// @Description makes task done
//// @Param id path string true "task id"
//// @Success 204
//// @Failure 404 "Task not found"
//// @Router /api/todo-list/tasks/{id}/done [put]
//func (c Controller) Done(ctx context.Context, w http.ResponseWriter, id string) {
//	err := c.s.Done(ctx, id)
//
//	if err != nil {
//		c.respondWithError(w, http.StatusNotFound, err.Error())
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
//
//// List принимает id извне так как для его получения используется специфичный метод роутера (сегодня chi, завтра echo)
//// @Summary list tasks
//// @Description lists tasks
//// @Param status query string  false "task status" Enums(active, done)
//// @Success 200 {object} dto.TaskResponse
//// @Router /api/todo-list/tasks [get]
//func (c Controller) List(ctx context.Context, w http.ResponseWriter, r *http.Request) {
//	statusParam := r.URL.Query().Get("status")
//
//	var status model.Status
//	if statusParam == "done" {
//		status = model.Done
//	} else {
//		status = model.BrandNew
//	}
//
//	tasks, err := c.s.List(ctx, status)
//	if err != nil {
//		c.respondWithError(w, http.StatusInternalServerError, "")
//		return
//	}
//
//	tasksResponse := make([]dto.TaskResponse, len(tasks))
//	for i, task := range tasks {
//		tasksResponse[i] = dto.TaskResponse{Id: task.Id, Title: task.Title, ActiveAt: dto.Datetime{Time: task.ActiveAt}}
//	}
//
//	json.NewEncoder(w).Encode(tasksResponse)
//}
//

func (c Controller) respondWithError(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.ErrorResponse{Message: errorMsg})
}
