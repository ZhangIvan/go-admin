package parameter

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Parameters struct {
	Page      string
	PageSize  string
	SortField string
	Columns   []string
	SortType  string
	Fields    map[string]string
}

var keys = []string{"__page", "__pageSize", "__sort", "__columns", "__prefix", "_pjax"}

const operatorSuffix = "__operator__"

func GetParam(values url.Values, defaultPageSize int, primaryKey, defaultSort string) Parameters {
	page := GetDefault(values, "__page", "1")
	pageSize := GetDefault(values, "__pageSize", strconv.Itoa(defaultPageSize))
	sortField := GetDefault(values, "__sort", primaryKey)
	sortType := GetDefault(values, "__sort_type", defaultSort)
	columns := GetDefault(values, "__columns", "")

	fields := make(map[string]string)

	for key, value := range values {
		if !modules.InArray(keys, key) && value[0] != "" {
			if key == "__sort_type" {
				if value[0] != "desc" && value[0] != "asc" {
					fields[key] = "desc"
				}
			} else {
				if strings.Contains(key, operatorSuffix) &&
					fields[strings.Replace(key, operatorSuffix, "", -1)] == "" {
					continue
				}
				fields[key] = value[0]
			}
		}
	}

	columnsArr := make([]string, 0)
	if columns != "" {
		columnsArr = strings.Split(columns, ",")
	}

	return Parameters{
		Page:      page,
		PageSize:  pageSize,
		SortField: sortField,
		SortType:  sortType,
		Fields:    fields,
		Columns:   columnsArr,
	}
}

func (param Parameters) GetFieldValue(field string) string {
	return param.Fields[field]
}

func (param Parameters) GetFieldOperator(field string) string {
	if param.Fields[field+operatorSuffix] == "" {
		return "="
	}
	return param.Fields[field+operatorSuffix]
}

func GetParamFromUrl(value string, fromList bool, defaultPageSize int, primaryKey, defaultSort string) Parameters {

	if !fromList {
		return Parameters{}
	}

	prevUrlArr := strings.Split(value, "?")
	paramArr := strings.Split(prevUrlArr[1], "&")

	var (
		page      = "1"
		pageSize  = strconv.Itoa(defaultPageSize)
		sortField = primaryKey
		sortType  = defaultSort
		columns   = make([]string, 0)
	)

	for i := 0; i < len(paramArr); i++ {
		arr := strings.Split(paramArr[i], "=")
		switch arr[0] {
		case "__pageSize":
			pageSize = arr[1]
		case "__page":
			page = arr[1]
		case "__sort":
			sortField = arr[1]
		case "__sort_type":
			sortType = arr[1]
		case "__columns":
			columns = strings.Split(arr[1], ",")
		}
	}

	return Parameters{
		Page:      page,
		PageSize:  pageSize,
		SortField: sortField,
		SortType:  sortType,
		Columns:   columns,
	}
}

func (param Parameters) SetPage(page string) Parameters {
	param.Page = page
	return param
}

func (param Parameters) GetRouteParamStr() string {
	return "?__page=" + param.Page + param.GetFixedParamStr()
}

func (param Parameters) GetRouteParamStrWithoutId() string {
	return regexp.MustCompile(`&id=[0-9]+`).ReplaceAllString(param.GetRouteParamStr(), "")
}

func (param Parameters) GetRouteParamStrWithoutPageSize() string {
	return "?__page=" + param.Page + param.GetFixedParamStrWithoutPageSize()
}

func (param Parameters) GetLastPageRouteParamStr() string {
	pageInt, _ := strconv.Atoi(param.Page)
	return "?__page=" + strconv.Itoa(pageInt-1) + param.GetFixedParamStr()
}

func (param Parameters) GetNextPageRouteParamStr() string {
	pageInt, _ := strconv.Atoi(param.Page)
	return "?__page=" + strconv.Itoa(pageInt+1) + param.GetFixedParamStr()
}

func (param Parameters) GetFixedParamStrWithoutPageSize() string {
	str := "&"
	for key, value := range param.Fields {
		str += key + "=" + value + "&"
	}
	if len(param.Columns) > 0 {
		return "&__columns=" + strings.Join(param.Columns, ",") + "&__sort=" + param.SortField +
			"&__sort_type=" + param.SortType + str[:len(str)-1]
	} else {
		return "&__sort=" + param.SortField + "&__sort_type=" + param.SortType + str[:len(str)-1]
	}
}

func (param Parameters) GetFixedParamStr() string {
	str := "&"
	for key, value := range param.Fields {
		str += key + "=" + value + "&"
	}
	if len(param.Columns) > 0 {
		return "&__columns=" + strings.Join(param.Columns, ",") + "&__pageSize=" + param.PageSize + "&__sort=" +
			param.SortField + "&__sort_type=" + param.SortType + str[:len(str)-1]
	} else {
		return "&__pageSize=" + param.PageSize + "&__sort=" +
			param.SortField + "&__sort_type=" + param.SortType + str[:len(str)-1]
	}
}

func GetDefault(values url.Values, key, def string) string {
	value := values.Get(key)
	if value == "" {
		return def
	}
	return value
}
