package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var emp = []employee{
	{Id: 1, Name: "divya", Email: "ds@gmail.com", Phone: "12345"},
	{Id: 2, Name: "nitish", Email: "nv@gmail.com", Phone: "12345"},
	{Id: 3, Name: "ritu", Email: "rs@gmail.com", Phone: "12345"},
	{Id: 4, Name: "sanskar", Email: "sa@gmail.com", Phone: "12345"},
	{Id: 5, Name: "isra", Email: "ie@gmail.com", Phone: "12345"},
}

func getEmp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, emp)
}

func createEmp(c *gin.Context) {
	var newEmp employee

	if err := c.BindJSON(&newEmp); err != nil {
		return
	}

	emp = append(emp, newEmp)
	c.IndentedJSON(http.StatusCreated, newEmp)

}

func main() {

	router := gin.Default()
	router.GET("/Emp", getEmp)
	router.Run("localhost:8080")
	router.POST("/create", createEmp)

}

// http.HandleFunc("/", home)
// http.HandleFunc("/about", about)
// //http.HandleFunc("/cal", cal)

// http.ListenAndServe(":8080", nil)

// func cal(w http.ResponseWriter, r *http.Request) {
// 	additionresult := addval(2, 3)
// 	divisionResult, err := divide(10.4, 3.3)
// 	fmt.Fprintf(w, "2+3 = %d", additionresult)
// 	fmt.Fprintf(w, "   ")
// 	if divisionResult == 0.0 {
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, "10.4/3.3 = %f", divisionResult)

// }

// func addval(x, y int) int {
// 	return x + y
// }

// func divide(x, y float32) (float32, error) {

// 	if y > x {
// 		err := errors.New("denominator is greater")
// 		return 0.0, err
// 	}
// 	result := x / y

// 	return result, nil
// }
