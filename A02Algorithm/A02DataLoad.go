package main

import "container/list"

//数据封装
//{
//    "bizType": "video",
//    "operateType": "playback",
//    "data": "{\"sn\":1212581947449081856,\"success\":1,\"message\":\"执行成功\",\
//   "videoRecordList\":[
//  {\"beginTime\":1596521410000, \"endTime\":1596535210000},
// {\"beginTime\":1596520410000,\"endTime\":1596535200000}
//]}"
//}


func getPostBodyData (time *int64, sn int64, list *list.List){
	result := make(map[string]string)
	result["bizType"] = "video"
	result["operateType"] = "playback"
	result["data"] = "data"

	data := make(map[string]interface{})
	data["sn"] = sn
	data["success"] = 1
	data["message"] = "执行成功"
	//timeEle := make(map[string]int64)

	data["videoRecordList"] = time

}
