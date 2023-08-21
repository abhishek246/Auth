package authenticator

import (
	"auth/internal/redisclient"
	"context"
	"encoding/json"
	"fmt"
)

func Authenticate(userId string, token string, resource string, action string) bool {
	fmt.Printf("Params | userId: %s token: %s resource: %s action: %s\n", userId, token, resource, action)
	ctx := context.Background()
	data, err := redisclient.RedisClient.Get(ctx, userId).Result()
	if err != nil || data == "" {
		return false
	}

	fmt.Printf("Access Validation Token: %s Data:%s \n", token, data)
	if token != data {
		return false
	}

	// Check for wheather the user has permission for the resource.
	// accessKey := fmt.Sprintf("%s_access_permission", data)
	accessPermissionRawData, err := redisclient.RedisClient.Get(ctx, token).Result()
	if err != nil {
		return false
	}
	accessPermission := make(map[string]map[string]string)
	if err := json.Unmarshal([]byte(accessPermissionRawData), &accessPermission); err != nil {
		return false
	}

	if _, ok := accessPermission["Admin"]; ok {
		return true
	}

	fmt.Printf("AccessPermission for user_id: %s permissions: %s", userId, accessPermissionRawData)
	if val, ok := accessPermission[resource]; ok {
		fmt.Println("Resource Found in permission: ", resource)
		if _, nestedOk := val[action]; nestedOk {
			fmt.Printf("Permission Found for resource %s and action %s", resource, action)
			return true
		} else {
			return false
		}
	} else {
		return false
	}

	return true
}
