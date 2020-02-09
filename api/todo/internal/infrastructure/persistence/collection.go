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

func getGroupReference(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID}, "/")
}

func getBoardCollection(groupRef string) string {
	return strings.Join([]string{groupRef, BoardCollection}, "/")
}
