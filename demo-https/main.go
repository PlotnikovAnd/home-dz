package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Task: Create an HTTP server with in-memory storage.
// Requirements:
// - Three endpoints (use net/http, no frameworks):
// 1. POST /resource — create a new item, return its ID
// 2. POST /resource/:id/action — perform an action on an item
// 3. GET /resource/:id — retrieve item state
// - In-memory storage (map or slice, thread-safe)
// - Return JSON responses
// - Handle errors (404, 400)
// Constraints:
// - No external dependencies except stdlib
// - Keep it minimal — no validation beyond basics
// - Must compile and run
// Deliverable: Working server, test with curl.
func main() {

}
