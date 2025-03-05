package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/30Piraten/beerPong/config"
	"github.com/gorilla/mux"
	"github.com/permitio/permit-golang/pkg/enforcement"
)

type BallRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Action string `json:"action"`
	Target string `json:"target"`
}

// throwBallHandler handles the initial req -> ball thrown
func throwBallHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	log.Println("Raw request body: ", string(body)) // For debugging!

	var req BallRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.UserID == "" || req.Action == "" || req.Target == "" {
		http.Error(w, "Missing request fields", http.StatusBadRequest)
		return
	}

	// Check if Redis client is initialized
	if config.RedisClient == nil {
		log.Println("❌ Redis client is not initialized!")
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	log.Printf("Checking permission for user: %s, action: %s:, target: %s", req.UserID, req.Action, req.Target)

	ctx := context.Background()
	// Save the req state in redis
	redisKey := fmt.Sprintf("ball:%s", req.UserID)

	if err := config.RedisClient.Set(ctx, redisKey, req.Target, time.Minute*5).Err(); err != nil {
		log.Printf("❌ Redis error: %v", err)
		log.Printf("Redis SET: ball:%s ex 300", req.Target) // Added here!
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("✅ Redis SET: %s -> %s (TTL: 5m)", redisKey, req.Target)

	// Now forward to Kafka (Placeholder)
	log.Printf("Ball thrown by %s for %s -> Passing to Kafka\n", req.UserID, req.Target)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Ball thrown!",
	})
}

func cupHandler(w http.ResponseWriter, r *http.Request) {

	// Check if permit.io is initialized
	if config.PermitClient == nil {
		log.Println("❌ Permit.io is not initialized")
		http.Error(w, "Permit.io failed initialization", http.StatusInternalServerError)
	}

	vars := mux.Vars(r)
	cupID := vars["cup_id"]

	var req BallRequest

	// Check the user permission for this cup
	userID := enforcement.UserBuilder(req.UserID).WithRoles([]enforcement.AssignedRole{{Role: "user"}}).Build()
	resource := enforcement.ResourceBuilder(cupID).Build()
	action := enforcement.Action("beer") // used for log

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	allowed, err := config.PermitClient.Check(userID, "beer", resource)

	if err != nil {
		http.Error(w, "Permission check failed", http.StatusInternalServerError)
		return
	}

	if allowed {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Access granted!",
			"cup":     cupID,
		})
	} else {
		http.Error(w, "Access denied!", http.StatusForbidden)
	}

	log.Printf("Checking permission for permit.check: %s, action: %s", userID, action)

}
