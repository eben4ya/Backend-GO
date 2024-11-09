package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"book_store/config"
	"book_store/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Employee Handlers
var employeeCollection *mongo.Collection = config.GetCollection("employees")

func GetEmployees(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var employees []models.Employee

    cursor, err := employeeCollection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var employee models.Employee
        if err := cursor.Decode(&employee); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        employees = append(employees, employee)
    }

    json.NewEncoder(w).Encode(employees)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var employee models.Employee
    _ = json.NewDecoder(r.Body).Decode(&employee)
    employee.ID = primitive.NewObjectID()
    employee.JoinDate = time.Now()

    _, err := employeeCollection.InsertOne(context.Background(), employee)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(employee)
}
