package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExerciseStore struct {
	mock.Mock
}

func (m *MockExerciseStore) CreateExercise(e *models.Exercise) error {
	args := m.Called(e)
	return args.Error(0)
}

func (m *MockExerciseStore) GetAllExercises() ([]models.Exercise, error) {
	args := m.Called()
	return args.Get(0).([]models.Exercise), args.Error(1)
}

func (m *MockExerciseStore) GetExercisesByUserID(id int) ([]*models.Exercise, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.Exercise), args.Error(1)
}

func (m *MockExerciseStore) DeleteExerciseByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandleCreateExercise(t *testing.T) {
	mockStore := new(MockExerciseStore)
	handler := &handlers.ExerciseHandler{
		Store: mockStore,
	}

	server := server.NewAPIServer(":8080", handler, nil, nil, nil)
	router := server.NewRouter()

	reqBody := handlers.CreateExerciseRequest{
		Name: "Pull-up",
	}
	body, _ := json.Marshal(reqBody)

	mockStore.On("CreateExercise", mock.Anything).Return(nil)

	req, err := http.NewRequest("POST", "/exercise", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `"Created Exercise: Pull-up"`
	assert.JSONEq(t, expectedBody, rr.Body.String())

	// Assert that the mock's CreateExercise was called once
	mockStore.AssertExpectations(t)
}

func TestHandleGetExercisesByUserID(t *testing.T) {
	mockStore := new(MockExerciseStore)
	handler := &handlers.ExerciseHandler{
		Store: mockStore,
	}
	server := server.NewAPIServer(":8080", handler, nil, nil, nil)
	router := server.NewRouter()

	// Mock the store's response to GetExercisesByUserID
	exercises := []*models.Exercise{
		{Name: "Push-up", UserID: 1},
		{Name: "Squat", UserID: 1},
	}

	mockStore.On("GetExercisesByUserID", 1).Return(exercises, nil)

	req, err := http.NewRequest("GET", "/exercise/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `[{"id":0,"name":"Push-up","user_id":1},{"id":0, "name":"Squat","user_id":1}]`
	assert.JSONEq(t, expectedBody, rr.Body.String())

	mockStore.AssertExpectations(t)
}

func TestHandleGetExercise(t *testing.T) {
	mockStore := new(MockExerciseStore)
	handler := &handlers.ExerciseHandler{
		Store: mockStore,
	}
	server := server.NewAPIServer(":8080", handler, nil, nil, nil)
	router := server.NewRouter()

	exercises := []models.Exercise{
		{ID: 1, Name: "Push-up", UserID: 1},
		{ID: 2, Name: "Squat", UserID: 1},
	}

	mockStore.On("GetAllExercises").Return(exercises, nil)

	req, err := http.NewRequest("GET", "/exercise", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `[{"id":1,"name":"Push-up","user_id":1},{"id":2, "name":"Squat","user_id":1}]`
	assert.JSONEq(t, expectedBody, rr.Body.String())

	mockStore.AssertExpectations(t)
}

func TestHandleDeleteExercise(t *testing.T) {
	mockStore := new(MockExerciseStore)
	handler := &handlers.ExerciseHandler{
		Store: mockStore,
	}
	server := server.NewAPIServer(":8080", handler, nil, nil, nil)
	router := server.NewRouter()

	mockStore.On("DeleteExerciseByID", 1).Return(nil)

	reqBody := handlers.DeleteExerciseRequest{
		ID: "1",
	}

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("DELETE", "/exercise", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `"Deleted exercise"`
	assert.JSONEq(t, expectedBody, rr.Body.String())

	mockStore.AssertExpectations(t)
}
