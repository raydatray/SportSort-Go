package user

import (
	"net/http"
	"strconv"

	"github.com/raydatray/sportsort-go/db"
)

type FilterUserParams struct {
	Types []db.UserType
	SportCenterIDs []int32
	Deleted *bool
	Page int
	PageSize int
}


func ParseFilterUsersParams(r *http.Request) (FilterUserParams, error) {
	query := r.URL.Query()

	params := FilterUserParams{
		Page:     1,
		PageSize: 20,
	}

	// Parse user types
	if types := query["types"]; len(types) > 0 {
		params.Types = make([]db.UserType, 0, len(types))
		for _, t := range types {
			userType := db.UserType(t)
			if err := userType.Scan(t); err == nil {
				params.Types = append(params.Types, userType)
			}
		}
	}

	// Parse sport center IDs
	if centerIDs := query["sport_center_ids"]; len(centerIDs) > 0 {
		params.SportCenterIDs = make([]int32, 0, len(centerIDs))
		for _, id := range centerIDs {
			if centerID, err := strconv.Atoi(id); err == nil {
				params.SportCenterIDs = append(params.SportCenterIDs, int32(centerID))
			}
		}
	}

	// Parse deleted status
	if deletedStr := query.Get("deleted"); deletedStr != "" {
		deleted, err := strconv.ParseBool(deletedStr)
		if err == nil {
			params.Deleted = &deleted
		}
	}

	// Parse pagination
	if page := query.Get("page"); page != "" {
		if pageNum, err := strconv.Atoi(page); err == nil && pageNum > 0 {
			params.Page = pageNum
		}
	}

	if pageSize := query.Get("page_size"); pageSize != "" {
		if size, err := strconv.Atoi(pageSize); err == nil && size > 0 {
			params.PageSize = size
		}
	}

	return params, nil
}
