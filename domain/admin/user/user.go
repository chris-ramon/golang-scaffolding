package user

import (
	"net/http"

	"github.com/admin-golang/admin"
	"github.com/admin-golang/admin/dataloader"
	"github.com/admin-golang/admin/icon"
	"github.com/admin-golang/admin/navigation"
)

func NewList() admin.Pager {
	return admin.NewListPage(admin.ListPageConfig{
		PageConfig: admin.PageConfig{
			ID:   "Users",
			URL:  "/users",
			Type: admin.ListPage,
			Icon: icon.Icon{
				Type: icon.Inventory,
			},
			ToolbarEnabled: true,
		},
		Title: "Users",
		MainButton: &admin.MainButton{
			Label: "Add User",
			URL:   "/users/create",
		},
		DataLoader: dataloader.New(dataloader.Config{
			URL:    "/users",
			Method: http.MethodGet,
			SearchParams: &navigation.SearchParams{
				navigation.SearchParam{
					Key: "limit",
					Value: navigation.SearchParamValue{
						FromQueryParams: true,
						SearchParamKey:  "limit",
					},
				},
				navigation.SearchParam{
					Key: "page",
					Value: navigation.SearchParamValue{
						FromQueryParams: true,
						SearchParamKey:  "page",
					},
				},
				navigation.SearchParam{
					Key: "searchTerm",
					Value: navigation.SearchParamValue{
						FromQueryParams: true,
						SearchParamKey:  "searchTerm",
					},
				},
			},
			HeaderConfig: &dataloader.HeaderConfig{
				Key: "Authorization",
				ValueConfig: dataloader.HeaderValueConfig{
					Prefix:            "Bearer ",
					AppStateFieldPath: "currentUser.jwt",
				},
			},
		}),
		Pagination: &admin.PaginationConfig{
			RowsPerPage: 25,
		},
		ListRowConfig: &admin.ListRowConfig{
			DataRowFieldName: "id",
			ParamKey:         ":id",
			OnClick: &admin.OnListRowClick{
				RedirectURL: "/users/:id",
			},
		},
		SearchConfig: &admin.ListSearchConfig{
			InputID:          "search",
			InputPlaceholder: "Search...",
		},
	})
}
