package persistence

import "strings"

const (
	// UserCollection - UserCollection名
	UserCollection = "users"
	// GroupCollection - GroupCollection名
	GroupCollection = "groups"
	// BoardCollection - BoardCollection名
	BoardCollection = "boards"
	// BoardListCollection - BoardListCollection名
	BoardListCollection = "board_lists"
	// TaskCollection - TaskCollection名
	TaskCollection = "tasks"
)

func getUserReference(userID string) string {
	return strings.Join([]string{UserCollection, userID}, "/")
}

func getGroupReference(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID}, "/")
}

func getUserCollection() string {
	return UserCollection
}

func getGroupCollection() string {
	return GroupCollection
}

func getBoardCollection(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID, BoardCollection}, "/")
}

func getBoardListCollection(groupID string, boardID string) string {
	return strings.Join([]string{GroupCollection, groupID, BoardCollection, boardID, BoardListCollection}, "/")
}

func getTaskCollection() string {
	return TaskCollection
}
