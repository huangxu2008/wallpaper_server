package models

import (
	"sort"
	"wallpaper_server/dao"
)

type WallpaperMenu struct {
	MenuID      uint            `json:"menu_id"`
	Name        string          `json:"name"`
	Path        string          `json:"path"`
	ParentID    uint            `json:"parent_id"`
	SortOrder   uint            `json:"sort_order"`
	Description string          `json:"description"`
	SubData     []WallpaperMenu `json:"sub_data"`
}

// 通过实现 sort.Interface 来自定义排序
type bySortOrder []WallpaperMenu

// Len 返回切片的长度
func (a bySortOrder) Len() int {
	return len(a)
}

// Less 比较两个元素的排序顺序
func (a bySortOrder) Less(i, j int) bool {
	return a[i].SortOrder < a[j].SortOrder
}

// Swap 交换两个元素
func (a bySortOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 构建树形结构
func buildMenuTree(menus []WallpaperMenu) []WallpaperMenu {
	// 这里只写死两层菜单，如果出现第三层菜单，将无法识别

	// 第一遍遍历，将所有菜单划分为一级菜单项和二级菜单项
	var rootMenus []WallpaperMenu
	subMenu := make(map[uint][]WallpaperMenu)
	for i := range menus {
		if menus[i].ParentID == 0 {
			rootMenus = append(rootMenus, menus[i])
		} else {
			subMenu[menus[i].ParentID] = append(subMenu[menus[i].ParentID], menus[i])
		}
	}

	// 将二级菜单项附加到一级菜单项中
	for i := range rootMenus {
		rootMenus[i].SubData = subMenu[rootMenus[i].MenuID]
	}

	// 对一级菜单进行排序
	sort.Sort(bySortOrder(rootMenus))

	// 对二级菜单进行排序
	for i := range rootMenus {
		sort.Sort(bySortOrder(rootMenus[i].SubData))
	}

	return rootMenus
}

func (WallpaperMenu) TableName() string {
	return "menu_list"
}

func GetMenuList() ([]WallpaperMenu, string, error) {
	// 获取菜单项
	var menuItems []WallpaperMenu
	if err := dao.Db.Find(&menuItems).Error; err != nil {
		return menuItems, "Failed to get menu lists", err
	}

	// 构建树形结构
	tree := buildMenuTree(menuItems)
	return tree, "Success to get menu lists", nil
}
