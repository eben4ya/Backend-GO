package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"book_store/config"
	"book_store/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Book Handlers
var bookCollection *mongo.Collection = config.GetCollection("books")

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var books []models.Book

    cursor, err := bookCollection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var book models.Book
        if err := cursor.Decode(&book); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        books = append(books, book)
    }

    json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var book models.Book
    err := bookCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book models.Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = primitive.NewObjectID()

    _, err := bookCollection.InsertOne(context.Background(), book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var updatedBook models.Book
    _ = json.NewDecoder(r.Body).Decode(&updatedBook)

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"stock": updatedBook.Stock, "price": updatedBook.Price}}

    _, err := bookCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Book updated successfully")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    _, err := bookCollection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Book deleted successfully")
}
