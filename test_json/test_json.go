package test_json

import (
	"encoding/hex"
	"encoding/json"
	"log"
)

func TestJsonMarshal() {
	postData := make(map[string]interface{})
	postData["drive_id"] = "driveId"
	postData["parent_file_id"] = "parentFileId"
	postData["limit"] = 100
	postData["all"] = false
	postData["url_expire_sec"] = 1600
	postData["image_thumbnail_process"] = "image/resize,w_400/format,jpeg"
	postData["image_url_process"] = "image/resize,w_1920/format,jpeg"
	postData["video_thumbnail_process"] = "video/snapshot,t_0,f_jpg,ar_auto,w_300"
	postData["fields"] = "*"
	postData["order_by"] = "updated_at"
	postData["order_direction"] = "DESC"

	data, err := json.MarshalIndent(postData, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
	log.Println("\n" + hex.Dump(data))

	json.Unmarshal(data, &postData)

	data, err = json.MarshalIndent(postData, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
}
