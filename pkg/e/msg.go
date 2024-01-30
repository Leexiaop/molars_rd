package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	//	Product
	ERROR_GET_PRODUCTS_FAIL:   "获取产品失败",
	ERROR_NOT_EXIST_PRODUCT:   "不存在该产品",
	ERROR_COUNT_PRODUCTS_FAIL: "获取产品总数失败",
	ERROR_EXIST_PRODUCT:       "产品名称已经存在",
	INVALID_PRODUCT_PARAMS:    "无效产品参数",
	ERROR_DELETE_PRODUCT_FAIL: "删除产品失败",
	ERROR_ADD_PRODUCT_FAIL:    "添加产品失败",
	ERROR_EXIST_PRODUCT_FAIL:  "产品已存在",
	ERROR_EDIT_PRODUCT_FAIL:   "编辑产品失败",

	//	Record
	ERROR_EXIST_RECORD:       "记录已存在",
	ERROR_EXIST_RECORD_LIST:  "该条记录下还有未删除的记录",
	ERROR_NOT_EXIST_RECORD:   "不存在该条记录",
	ERROR_COUNT_RECORDS_FAIL: "获取记录总数失败",
	ERROR_GET_RECORDS_FAIL:   "获取记录失败",
	ERROR_EXPORT_RECORD_FAIL: "导出记录失败",
	ERROR_EXIST_RECORD_FAIL:  "记录已存在",
	ERROR_DELETE_RECORD_FAIL: "删除记录失败",
	ERROR_ADD_RECORD_FAIL:    "添加记录失败",
	ERROR_EDIT_RECORD_FAIL:   "编辑记录失败",

	//	user
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_GET_USER_FAIL:            "获取用户信息失败",
	ERROR_COUNT_USER_LIST_FAIL:     "获取用户列表失败",
	ERROR_EXIST_USER_FAIL:          "用户已存在",
	ERROR_NOT_EXIST_USER:           "不存在该用户",
	ERROR_ADD_USER_FAIL:            "添加用户失败",

	//	上传图片
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "上传图片失败",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:  "保存图片失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}