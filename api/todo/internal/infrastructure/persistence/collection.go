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

func GetGroupReference(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID}, "/")
}

func GetBoardCollection(groupID string) string {
	return strings.Join([]string{GroupCollection, groupID, BoardCollection}, "/")
}

func GetBoardListCollection(groupID string, boardID string) string {
	return strings.Join([]string{GroupCollection, groupID, BoardCollection, boardID, BoardListCollection}, "/")
}

func GetTaskCollection() string {
	return TaskCollection
}
