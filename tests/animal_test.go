package tests

import (
	"bytes"
	"github.com/MeGaNeKoS/TF-Backend/database"
	"github.com/MeGaNeKoS/TF-Backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testRouter *gin.Engine

func init() {
	route := gin.Default()
	routes.Setup(route)
	testRouter = route
	//os.Setenv("DB_NAME", "implementation")
	//os.Setenv("ENVIRONMENT", "test")

	memSql := sqlite.Open(":memory:")
	db, _ := gorm.Open(memSql)

	database.SetSqliteDB(db)

	db.Exec("DROP TABLE IF EXISTS animals")
	db.Exec(`
		CREATE TABLE IF NOT EXISTS animals (
		    			id INTEGER PRIMARY KEY AUTOINCREMENT,
		    			name TEXT NOT NULL UNIQUE,
		    			class TEXT NOT NULL,
		    			legs INT NOT NULL
		                                   );
		`)
}

/*func TestMockDB(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	jsonStr := []byte(`{
						"name": "cat",
						"class": "mammals",
						"legs":4
						}`)
	// add animal to database for the first time
	req1, err := http.NewRequest("POST", "/api/v1/add", bytes.NewBuffer(jsonStr))
	assert.Nil(t, err)
	resp1 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp1, req1)
	assert.Equal(t, 200, resp1.Code)

}*/

func TestAddAnimal(t *testing.T) {
	// TODO
	defer httpmock.DeactivateAndReset()

	jsonStr := []byte(`{
						"name": "cat",
						"class": "mammals",
						"legs":4
						}`)
	// add animal to database for the first time
	req1, err := http.NewRequest("POST", "/api/v1/add", bytes.NewBuffer(jsonStr))
	assert.Nil(t, err)
	resp1 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp1, req1)
	assert.Equal(t, 200, resp1.Code)

	// add animal to database for the second time
	req2, err := http.NewRequest("POST", "/api/v1/add", bytes.NewBuffer(jsonStr))
	assert.Nil(t, err)
	resp2 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp2, req2)
	assert.Equal(t, 400, resp2.Code)
}

func TestEditAnima(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	jsonStr1 := []byte(`{
						"id": 1,
						"name": "cat",
						"class": "mammals",
						"legs":4
						}`)
	// edit animal in database
	req1, err := http.NewRequest("PUT", "/api/v1/update", bytes.NewBuffer(jsonStr1))
	assert.Nil(t, err)
	resp1 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp1, req1)
	assert.Equal(t, 200, resp1.Code)

	// edit animal in database but not exist
	jsonStr2 := []byte(`{
						"id": 10,
						"name": "dog",
						"class": "birds",
						"legs":4
						}`)
	req2, err := http.NewRequest("PUT", "/api/v1/update", bytes.NewBuffer(jsonStr2))
	assert.Nil(t, err)
	resp2 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp2, req2)
	assert.Equal(t, 400, resp2.Code)
}

func TestGetAnimal(t *testing.T) {

	defer httpmock.DeactivateAndReset()

	res1, err := http.NewRequest("GET", "http://localhost:8080/api/v1/get", nil)
	assert.Nil(t, err)
	resp1 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp1, res1)
	assert.Equal(t, 200, resp1.Code)

	req2, err := http.NewRequest("GET", "http://localhost:8080/api/v1/get/1", nil)
	assert.Nil(t, err)
	resp2 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp2, req2)
	assert.Equal(t, 200, resp2.Code)

	res3, err := http.NewRequest("GET", "http://localhost:8080/api/v1/get/50", nil)
	assert.Nil(t, err)
	resp3 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp3, res3)
	assert.Equal(t, 404, resp3.Code)
}

func TestDeleteAnimal(t *testing.T) {

	defer httpmock.DeactivateAndReset()

	jsonStr1 := []byte(`{
						"id": 1,
						"name": "cat",
						"class": "mammals",
						"legs":4
						}`)

	// delete animal in database
	req1, err := http.NewRequest("DELETE", "/api/v1/delete", bytes.NewBuffer(jsonStr1))
	assert.Nil(t, err)
	resp1 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp1, req1)
	assert.Equal(t, 204, resp1.Code)

	// delete animal in database but not exist
	jsonStr2 := []byte(`{
						"id": 10,
						"name": "dog",
						"class": "birds",
						"legs":4
						}`)
	req2, err := http.NewRequest("DELETE", "/api/v1/delete", bytes.NewBuffer(jsonStr2))
	assert.Nil(t, err)
	resp2 := httptest.NewRecorder()
	testRouter.ServeHTTP(resp2, req2)
	assert.Equal(t, 404, resp2.Code)
}
