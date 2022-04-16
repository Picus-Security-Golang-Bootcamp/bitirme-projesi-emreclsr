package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	Sort         string      `json:"sort"`
	TotalRows    int         `json:"total_rows"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
	FromRow      int         `json:"from_row"`
	ToRow        int         `json:"to_row"`
	Rows         interface{} `json:"rows"`
	TotalPages   int         `json:"total_pages"`
	URLPath      string      `json:"url_path"`
	Searchs      []Search    `json:"searchs"`
}

type Search struct {
	Column string `json:"column"`
	Action string `json:"action"`
	Query  string `json:"query"`
}

func GeneratePaginationRequest(context *gin.Context) *Pagination {
	// default limit, page & sort parameter
	limit := 10
	page := 1
	sort := "created_at desc"

	var searchs []Search

	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}

		// check if query parameter key contains dot
		if strings.Contains(key, ".") {
			// split query parameter key by dot
			searchKeys := strings.Split(key, ".")

			// create search object
			search := Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}
