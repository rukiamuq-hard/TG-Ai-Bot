package models

type History struct {
	Name string `bson:"name"`
	UID  int64  `bson:"uid"`
	Text string `bson:"text"`
	CID  int64  `bson:"cid"`
	MID  int    `bson:"mid"`
}
