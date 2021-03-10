package practice_1

func destCity(paths [][]string) string {
	//map的初始化用make比较好，用var不赋值并不能初始化
	mapAB := make(map[string]string, 1)
	//注意这是一个二维数组，用hash去存入所有起点和终点
	for _, val := range paths {
		mapAB[val[0]] = val[1]
	}
	//筛选结果
	for _, val := range paths {
		_, exist := mapAB[val[1]]
		if !exist {
			return val[1]
		}
	}
	//语法必要，如果题没问题，不会走到这儿
	return ""
}
