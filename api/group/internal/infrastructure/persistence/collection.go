package persistence

import "strings"

const (
	// UserCollection - UserCollection名
	UserCollection = "users"
	// GroupCollection - GroupCollection名
	GroupCollection = "groups"
	// BoardCollection - BoardCollection名
	BoardCollection = "boards"
)

func GetUserReference(userID string) string {
	return strings.Join([]string{UserCollection, userID}, "/")
}

func GetGroupID(groupRef string) string {
	slice := strings.Split(groupRef, "/")
	return slice[len(slice)-1]
}

func GetGroupReference(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID}, "/")
}

func GetBoardCollection(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID, BoardCollection}, "/")
}
