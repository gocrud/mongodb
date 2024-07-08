// 用于业务中的也一些扩展方法

package mongodb

const bizPageListCountField = "__biz_page_ls_count__"

// BizPageList 业务查询方法，分页查询列表
//   - page 页码，最小值为1
//   - size 每页数据条数，小于等于0时，该值默认为10
//   - data 查询数据变量的指针值
//   - 返回值(符合条件的总数,错误)
func (a *Aggregation) BizPageList(page int64, size int64, data any) (int64, error) {
	if page < 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	aggsCount := a.Clone()
	aggsCount.Count(bizPageListCountField)
	var countData []map[string]int64
	var err error
	err = aggsCount.FindMany(&countData)
	if err != nil {
		return 0, err
	}
	if countData != nil {
		_count := countData[0][bizPageListCountField]
		if _count <= 0 {
			return 0, nil
		}
		a.Skip((page - 1) * size)
		a.Limit(size)
		err = a.FindMany(data)
		if err != nil {
			return 0, err
		}
		return _count, nil
	}

	return 0, nil
}
