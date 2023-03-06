package service

// service 通用方法

// 分页, 排序处理
func parsePageAndSort(pageNumber, pageSize int, isAsc bool) (skipNum, sort int) {
	skipNum = (pageNumber - 1) * pageSize

	if isAsc {
		sort = 1
	} else {
		sort = -1
	}
	return
}
