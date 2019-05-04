package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/SevereCloud/vksdk/5.92/object"
)

// LikesAddResponse struct
type LikesAddResponse struct {
	Likes int `json:"likes"`
}

// LikesAdd adds the specified object to the Likes list of the current user.
// https://vk.com/dev/likes.add
func (vk VK) LikesAdd(params map[string]string) (response LikesAddResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.add", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesDeleteResponse struct
type LikesDeleteResponse struct {
	Likes int `json:"likes"`
}

// LikesDelete deletes the specified object from the Likes list of the current user.
// https://vk.com/dev/likes.delete
func (vk VK) LikesDelete(params map[string]string) (response LikesDeleteResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.delete", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesGetListResponse struct
type LikesGetListResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// LikesGetList likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
// https://vk.com/dev/likes.getList
func (vk VK) LikesGetList(params map[string]string) (response LikesGetListResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("likes.getList", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesGetListExtendedResponse struct
type LikesGetListExtendedResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// LikesGetListExtended likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
// https://vk.com/dev/likes.getList
func (vk VK) LikesGetListExtended(params map[string]string) (response LikesGetListExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("likes.getList", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesIsLikedResponse struct
type LikesIsLikedResponse struct {
	Liked  int `json:"liked"`
	Copied int `json:"copied"`
}

// LikesIsLiked checks for the object in the Likes list of the specified user.
// https://vk.com/dev/likes.isLiked
func (vk VK) LikesIsLiked(params map[string]string) (response LikesIsLikedResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.isLiked", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
