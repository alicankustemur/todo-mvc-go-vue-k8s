package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	createBuySomeMilkTodo = `{"title":"buy some milk","completed":false}`
	createEnjoyTheAssignment = `{"title":"enjoy the assignment","completed":false}`
	buySomeMilkTodoResponse = `{"id":1,"title":"buy some milk","order":0,"completed":false,"url":"http://example.com/todos/1"}`
	todoNotFoundResponse = `Todo note was not found`
	emptyResponse = ``
)

func TestGetAllTodos(t *testing.T) {
	resetDatabaseBefore()
	
	getAllTodosResponse := `[{"id":1,"title":"buy some milk","order":0,"completed":false,"url":"http://example.com/todos/1"},{"id":2,"title":"enjoy the assignment","order":0,"completed":false,"url":"http://example.com/todos/2"}]`

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	createTodoHandler(c1)
	
	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createEnjoyTheAssignment))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	createTodoHandler(c2)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// Assertions
	if assert.NoError(t, getAllTodosHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getAllTodosResponse, filterNewLines(rec.Body.String()))
	}
}

func TestCreateTodo(t *testing.T) {
	resetDatabaseBefore()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, createTodoHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, buySomeMilkTodoResponse, filterNewLines(rec.Body.String()))
	}
}

func TestCreateTodoError(t *testing.T) {
	resetDatabaseBefore()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/todos", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, createTodoHandler(c))
}

func TestDeleteAllTodos(t *testing.T) {
	resetDatabaseBefore()

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	createTodoHandler(c1)
	
	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createEnjoyTheAssignment))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	createTodoHandler(c2)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// Assertions
	if assert.NoError(t, deleteAllTodosHandler(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, emptyResponse, filterNewLines(rec.Body.String()))
	}
}

func TestGetTodo(t *testing.T) {
	resetDatabaseBefore()

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	createTodoHandler(c1)
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, getTodoHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, buySomeMilkTodoResponse, filterNewLines(rec.Body.String()))
	}
}

func TestGetTodoError(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	
	// Assertions
	assert.Error(t, getTodoHandler(c))
}

func TestGetTodoNotFound(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	
	// Assertions
	if assert.NoError(t, getTodoHandler(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, todoNotFoundResponse, filterNewLines(rec.Body.String()))
		assert.NotEqual(t, buySomeMilkTodoResponse, filterNewLines(rec.Body.String()))
	}
}

func TestDeleteTodo(t *testing.T) {
	resetDatabaseBefore()

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	createTodoHandler(c1)
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, deleteTodoHandler(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, emptyResponse, filterNewLines(rec.Body.String()))
	}
}

func TestDeleteTodoError(t *testing.T) {
	resetDatabaseBefore()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	
	// Assertions
	assert.Error(t, deleteTodoHandler(c)) 
}

func TestUpdateTodo(t *testing.T) {
	resetDatabaseBefore()

	updatedTodoJSON := `{"id":1,"title":"enjoy the assignment","order":0,"completed":false,"url":"http://example.com/todos/1"}`

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(createBuySomeMilkTodo))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	createTodoHandler(c1)
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(updatedTodoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, updateTodoHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, updatedTodoJSON, filterNewLines(rec.Body.String()))
	}

	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(updatedTodoJSON))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)
	c2.SetPath("/todos/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("3")

	// Assertions
	if assert.NoError(t, updateTodoHandler(c2)) {
		assert.Equal(t, http.StatusNotFound, rec2.Code)
		assert.NotEqual(t, updatedTodoJSON, filterNewLines(rec2.Body.String()))
		assert.Equal(t, todoNotFoundResponse, filterNewLines(rec2.Body.String()))
	}
}

func TestUpdateTodoError(t *testing.T) {
	resetDatabaseBefore()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(createBuySomeMilkTodo))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")

	// Assertions
	assert.Error(t, updateTodoHandler(c))
}

func resetDatabaseBefore() {
	newInMemoryTodoRepository()
}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}